CREATE TABLE IF NOT EXISTS words
(
    id   INTEGER PRIMARY KEY NOT NULL,
    name TEXT UNIQUE         NOT NULL
);

CREATE TABLE IF NOT EXISTS similar_words
(
    word_id         INTEGER NOT NULL,
    word_similar_id INTEGER NOT NULL,
    similarity      REAL    NOT NULL,
    PRIMARY KEY (word_id, word_similar_id),
    FOREIGN KEY (word_id) REFERENCES words (id),
    FOREIGN KEY (word_similar_id) REFERENCES words (id)
);