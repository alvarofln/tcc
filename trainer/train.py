import csv
import os
import re
import sqlite3
from collections import Counter
from os import path
from ordered_set import OrderedSet

from gensim.models import Word2Vec
from sklearn.decomposition import PCA
from matplotlib import pyplot as plt


import logging
import json

from spylls.hunspell import Dictionary

PLB_WORDS_FILE = "data/plb-words.csv"
LOOKUP_CACHE_FILE = "data/lookup_cache.json"
LYRICS_FILE = "data/lyrics.txt"
WORD2VEC_MODEL_FILE = "data/word2vec.model"
WORD2VEC_JSON_MODEL_FILE = "data/word2vec.json"
DATABASE_FILE = "data/word2vec.db"
DICTIONARY_FILE = "data/pt_BR"

logging.basicConfig(format="%(asctime)s : %(levelname)s : %(message)s", level=logging.INFO)

dictionary = Dictionary.from_files(DICTIONARY_FILE)


def load_lookup_cache():
    if path.exists(LOOKUP_CACHE_FILE):
        with open(LOOKUP_CACHE_FILE, "r") as infile:
            return json.load(infile)
    return dict()


lookup = load_lookup_cache()


def save_lookup_cache():
    if not os.path.exists(LOOKUP_CACHE_FILE):
        with open(LOOKUP_CACHE_FILE, "w") as outfile:
            json.dump(lookup, outfile, indent=4)


def tokenize(text):
    # Expressão regular para encontrar palavras no texto
    pattern = r'\b\w+\b'
    return re.findall(pattern, text.lower())


def dump_top_1000_words():
    plbDict = {}

    # carrega dicionario de palavras do portal da lingua portuguesa
    with open(PLB_WORDS_FILE, 'r', encoding='utf-8') as csvfile:
        reader = csv.reader(csvfile)
        for row in reader:
            if row:
                key = row[0]
                value = row[1]
                plbDict[key] = value

    with open(LYRICS_FILE, 'r', encoding='utf-8') as file:
        text = file.read()
    words = tokenize(text)
    word_count = Counter(words)
    word_count = {k: v for k, v in word_count.items() if k in plbDict}
    top_words = dict(sorted(word_count.items(), key=lambda item: item[1], reverse=True)[:1000])

    with open('top_words.csv', 'w', newline='', encoding='utf-8') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow(['Palavra', 'Frequência', 'Separação'])  # Cabeçalho do CSV
        for word, frequency in top_words.items():
            writer.writerow([word, frequency, plbDict[word]])


def lookup_word(word: str):
    if word not in lookup:
        lookup[word] = dictionary.lookup(word)
    return lookup[word]


def create_model():
    with open(LYRICS_FILE, "r", encoding="utf-8") as infile:
        lines = infile.readlines()
        # Preparando os dados para o Word2Vec
        # Neste caso, estamos tratando cada linha como uma "frase" ou sequência de palavras.
        sentences = [list(OrderedSet(tokenize(line))) for line in lines if line.strip() != ""]
        # filtra sentenças
        sentences = [sentence for sentence in sentences if
                     len(sentence) >= 3 and all([lookup_word(word) for word in sentence])]
        save_lookup_cache()
        model = Word2Vec(sentences, vector_size=100, window=5, min_count=2, workers=8)
        model.train(sentences, total_examples=len(sentences), epochs=100)
        model.save(WORD2VEC_MODEL_FILE)
        return model


def load_model():
    if not path.exists(WORD2VEC_MODEL_FILE):
        return create_model()
    return Word2Vec.load(WORD2VEC_MODEL_FILE)


def create_db():
    conn = sqlite3.connect(DATABASE_FILE)
    c = conn.cursor()
    # Criar a tabela words
    c.execute('''
    CREATE TABLE IF NOT EXISTS words (
        id INTEGER PRIMARY KEY NOT NULL,
        name TEXT UNIQUE NOT NULL
    )
    ''')

    # Criar a tabela similar_words
    c.execute('''
    CREATE TABLE IF NOT EXISTS similar_words (
        word_id INTEGER NOT NULL,
        word_similar_id INTEGER NOT NULL,
        similarity REAL NOT NULL,
        PRIMARY KEY (word_id, word_similar_id),
        FOREIGN KEY (word_id) REFERENCES words(id),
        FOREIGN KEY (word_similar_id) REFERENCES words(id)
    )
    ''')
    conn.commit()
    conn.close()


def save_words_db(model):
    conn = sqlite3.connect(DATABASE_FILE)
    c = conn.cursor()
    for word in model.wv.index_to_key:
        word_index = model.wv.key_to_index[word]
        c.execute('INSERT INTO words (id, name) VALUES (?, ?)', (word_index, word))
    conn.commit()
    conn.close()


def save_similar_words_db(model):
    conn = sqlite3.connect(DATABASE_FILE)
    c = conn.cursor()
    for word in model.wv.index_to_key:
        word_id = model.wv.key_to_index[word]
        similar_words = model.wv.most_similar(word, topn=250)
        for similar_word in similar_words:
            similar_id = model.wv.key_to_index[similar_word[0]]
            similarity = similar_word[1]
            c.execute('INSERT INTO similar_words (word_id, word_similar_id, similarity) VALUES (?, ?, ?)',
                      (word_id, similar_id, similarity))
    conn.commit()
    conn.close()


def vacum_db():
    conn = sqlite3.connect(DATABASE_FILE)
    c = conn.cursor()
    c.execute('VACUUM')
    conn.commit()
    conn.close()


def save_model_db(model):
    create_db()
    save_words_db(model)
    save_similar_words_db(model)
    vacum_db()


def load_similar_words_db(word: str):
    conn = sqlite3.connect(DATABASE_FILE)
    c = conn.cursor()
    c.execute('''
    SELECT w.name, sw.similarity
    FROM similar_words sw
    INNER JOIN words w ON sw.word_similar_id = w.id
    WHERE sw.word_id = (SELECT id FROM words WHERE name = ?)
    ORDER BY similarity DESC
    ''', (word,))
    result = c.fetchall()
    conn.close()
    return result

def train_small_example():
    # Exemplo simples de corpus
    sentences = [
        ["cat", "say", "meow"],
        ["dog", "say", "woof"],
        ["cow", "say", "moo"],
    ]

    # Treinando o modelo Word2Vec
    model = Word2Vec(sentences, vector_size=100, window=5, min_count=1, workers=4)

    # Palavras para visualização
    words = list(model.wv.index_to_key)
    vectors = [model.wv[word] for word in words]

    # Reduz a dimensionalidade para 2D usando PCA
    pca = PCA(n_components=2)
    vectors2d = pca.fit_transform(vectors)

    # Cria um gráfico com as palavras no espaço 2D
    plt.figure(figsize=(10, 8))
    plt.scatter(vectors2d[:, 0], vectors2d[:, 1], edgecolors='k', c='r')
    for word, (x, y) in zip(words, vectors2d):
        plt.text(x, y, word, fontsize=9)
    #plt.title("Visualização 2D das Palavras com Word2Vec")
    #plt.xlabel("Componente PCA 1")
    #plt.ylabel("Componente PCA 2")
    plt.grid(True)
    plt.show()

if __name__ == '__main__':
    model = load_model()
    save_model_db(model)
    dump_top_1000_words()
    train_small_example()
    print("done.")
