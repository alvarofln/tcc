package word

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestTopWords(t *testing.T) {
	file, err := os.Open("top_words.csv") //crawled from http://www.portaldalinguaportuguesa.org
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil { //skip headers
		panic(err)
	}

	for {
		r, err := reader.Read()
		if err != nil {
			break // Stop at EOF
		}

		syllables := Syllabificate(r[0])
		stressType := DetectStress(syllables)
		w := NewWord(-1, r[0], syllables, stressType)

		if w.String() != r[2] {
			t.Errorf("PLB expected %s, we got %s", r[2], w.String())
		}
	}
}
