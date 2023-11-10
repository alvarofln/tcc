package lyric

import (
	"context"
	"gotoolapi/word"
	"strings"
)

type Service struct {
	ws *word.Service
}

func NewService(ws *word.Service) *Service {
	return &Service{ws: ws}
}

func (s *Service) Process(ctx context.Context, input *InputLyric) (*Lyric, error) {
	lyricTextInput := input.Title + "\n" + input.Body
	textLines := strings.Split(lyricTextInput, "\n")

	lyric := new(Lyric)
	var lyricLines []*LyricLine

	for _, textLine := range textLines {
		syllableCount := 0
		fields := strings.Fields(textLine)

		var words []*word.Word

		for _, field := range fields {
			if len(strings.Trim(field, " ")) == 0 {
				continue
			}
			w := s.findWordByName(ctx, field)

			if len(w.Syllables) == 0 {
				continue
			}

			words = append(words, w)
			syllableCount += len(w.Syllables)
		}
		if len(words) == 0 {
			continue
		}
		lyricLines = append(lyricLines, &LyricLine{words, syllableCount})
	}

	lyric.Lines = lyricLines

	return lyric, nil
}

func (s *Service) findWordByName(ctx context.Context, name string) *word.Word {
	w, err := s.ws.FindWordByName(ctx, strings.ToLower(name))

	if err != nil {
		w = new(word.Word)
		w.ID = -1
	}

	w.Name = name
	w.Syllables = word.Syllabificate(w.Name)
	w.StressType = word.DetectStress(w.Syllables)

	return w
}
