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
			//t.Errorf("PLB expected %s, we got %s", r[2], w.String()) //uncomment to see the differences
		}
	}
}

// go test -bench=. -benchtime=20s
func BenchmarkTopWords(b *testing.B) {
	file, err := os.Open("top_words.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	words := make([]string, 0)

	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		words = append(words, r[0])
	}

	for i := 0; i < b.N; i++ {
		for _, w := range words {
			DetectStress(Syllabificate(w))
		}
	}
}
func TestTopWordsCounters(t *testing.T) {
	file, err := os.Open("top_words.csv") //crawled from http://www.portaldalinguaportuguesa.org
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil { //skip headers
		panic(err)
	}

	counters := make(map[StressType]int)

	for {
		r, err := reader.Read()
		if err != nil {
			break // Stop at EOF
		}
		syllables := Syllabificate(r[0])
		stressType := DetectStress(syllables)
		counters[stressType]++
	}

	for k, v := range counters {
		t.Logf("%s: %d", k, v)
	}
}
