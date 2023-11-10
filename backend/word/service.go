package word

import (
	"context"
	"gotoolapi/word/db"
	"sort"
	"strings"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q: q}
}

func (s *Service) FindWordById(ctx context.Context, ID int64) (*Word, error) {
	w, err := s.q.FindWordById(ctx, ID)

	if err != nil {
		return nil, err
	}

	return &Word{
		ID:   w.ID,
		Name: w.Name,
	}, nil
}

func (s *Service) FindWordByName(ctx context.Context, name string) (*Word, error) {
	w, err := s.q.FindWordByName(ctx, name)

	if err != nil {
		return nil, err
	}

	return &Word{
		ID:   w.ID,
		Name: w.Name,
	}, nil
}

func (s *Service) FindSimilarWordsById(ctx context.Context, ID int64) ([]*SimilarWord, error) {
	sws, err := s.q.FindSimilarWordsById(ctx, ID)

	if err != nil {
		return nil, err
	}

	var similarWords []*SimilarWord

	for _, sw := range sws {
		similarWords = append(similarWords, &SimilarWord{
			Word: Word{
				ID:   sw.ID,
				Name: sw.Name,
			},
			Similarity: sw.Similarity,
		})
	}

	return similarWords, nil
}

func (s *Service) FindAllWords(ctx context.Context) ([]*Word, error) {
	words, err := s.q.FindAllWords(ctx)

	if err != nil {
		return nil, err
	}

	var result []*Word

	for _, w := range words {
		result = append(result, s.createWord(&w))
	}

	return result, nil
}

func (s *Service) createWord(w *db.Word) *Word {
	syllables := Syllabificate(w.Name)
	stressType := DetectStress(syllables)
	return NewWord(w.ID, w.Name, syllables, stressType)
}
func (s *Service) FindRhymingWords(ctx context.Context, word string) ([]*Word, error) {
	targetWord := s.createWord(&db.Word{ID: -1, Name: strings.ToLower(word)})

	words, _ := s.FindAllWords(ctx)

	var result []*Word

	for _, w := range words {
		if w.StressType != targetWord.StressType || targetWord.Name == w.Name {
			continue
		}
		result = append(result, w)
	}
	targetSyllablesLen := len(targetWord.Syllables)
	sort.Slice(result, func(i, j int) bool {
		return Levenshtein(targetWord.Name, result[i].Name) < Levenshtein(targetWord.Name, result[j].Name) && targetSyllablesLen == len(result[i].Syllables)
	})

	if len(result) > 4000 {
		result = result[:4000]
	}

	return result, nil
}
