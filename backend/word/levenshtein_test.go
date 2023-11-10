package word

import "testing"

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		wordA    string
		wordB    string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"Alvaro", "Alvaro", 0},
		{"Alvaro", "Auvaro", 1},
		{"Alvaro", "Álvaro", 1},
		{"Alvaro", "alvaro", 1},
		{"Alvaro", "Alvaros", 1},
		{"Alvaro", "Alvarós", 2},
		{"Alvaro", "", 6},
	}

	for _, test := range tests {
		distance := Levenshtein(test.wordA, test.wordB)
		if distance != test.expected {
			t.Errorf("Levenshtein(%s, %s) = %d; expected %d", test.wordA, test.wordB, distance, test.expected)
		}
	}
}
