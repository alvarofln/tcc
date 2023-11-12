package word

import (
	"testing"
)

func TestSyllabificate(t *testing.T) {
	tests := []struct {
		word     string
		expected []string
	}{
		{"água", []string{"á", "gua"}},
		{"amar", []string{"a", "mar"}},
		{"amaste", []string{"a", "mas", "te"}},
		{"amarei", []string{"a", "ma", "rei"}},
		{"amaria", []string{"a", "ma", "ri", "a"}},
		{"correria", []string{"cor", "re", "ri", "a"}},
		{"ressurreição", []string{"res", "sur", "rei", "ção"}},
		{"inconstitucionalissimamente", []string{"in", "cons", "ti", "tu", "ci", "o", "na", "lis", "si", "ma", "men", "te"}},
		{"canções", []string{"can", "ções"}},
		{"aéreo", []string{"a", "é", "re", "o"}},
		{"Álvaro", []string{"ál", "va", "ro"}},
		{"último", []string{"úl", "ti", "mo"}},
		{"série", []string{"sé", "ri", "e"}},
		{"história", []string{"his", "tó", "ri", "a"}},
		{"cãibra", []string{"cãi", "bra"}},
		{"bombom", []string{"bom", "bom"}},
		{"saúva", []string{"sa", "ú", "va"}},
		{"Abigail", []string{"a", "bi", "ga", "il"}},
		{"ruiu", []string{"ru", "iu"}},
		{"Rui", []string{"rui"}},
		{"ruim", []string{"ru", "im"}},
		{"ideia", []string{"i", "dei", "a"}},
		{"ideias", []string{"i", "dei", "as"}},
		{"casa", []string{"ca", "sa"}},
		{"uruguaio", []string{"u", "ru", "guai", "o"}},
		{"rainha", []string{"ra", "i", "nha"}},
		{"rinha", []string{"ri", "nha"}},
		{"canções", []string{"can", "ções"}},
		{"avião", []string{"a", "vi", "ão"}},
		{"mãe", []string{"mãe"}},
		{"mães", []string{"mães"}},
		{"paraguai", []string{"pa", "ra", "guai"}},
		{"acre", []string{"a", "cre"}},
		{"algum", []string{"al", "gum"}},
		{"andam", []string{"an", "dam"}},
		{"uivo", []string{"ui", "vo"}},
		{"Caetano", []string{"ca", "e", "ta", "no"}},
		{"pais", []string{"pais"}},
		{"país", []string{"pa", "ís"}},
		{"ideinha", []string{"i", "de", "i", "nha"}},
		{"caos", []string{"ca", "os"}},
		{"pérola", []string{"pé", "ro", "la"}},
		{"guaraná", []string{"gua", "ra", "ná"}},
		{"mulherão", []string{"mu", "lhe", "rão"}},
		{"ideinha", []string{"i", "de", "i", "nha"}},
		{"rainha", []string{"ra", "i", "nha"}},
		{"independente", []string{"in", "de", "pen", "den", "te"}},
		{"joia", []string{"joi", "a"}},
		{"caiu", []string{"ca", "iu"}},
		{"uva", []string{"u", "va"}},
		{"pai", []string{"pai"}},
		{"país", []string{"pa", "ís"}},
		{"saara", []string{"sa", "a", "ra"}},
		{"aguado", []string{"a", "gua", "do"}},
		{"guaraná", []string{"gua", "ra", "ná"}},
		{"oxítona", []string{"o", "xí", "to", "na"}},
		{"Piauí", []string{"pi", "au", "í"}},
		{"ao", []string{"ao"}},
	}

	for _, test := range tests {
		syllables := Syllabificate(test.word)

		if len(syllables) != len(test.expected) {
			t.Errorf("Syllabificate(%v) = %v; want %v", test.word, syllables, test.expected)
		}

		for i := range syllables {
			if syllables[i] != test.expected[i] {
				t.Errorf("Syllabificate(%v) = %v; want %v", test.word, syllables, test.expected)
			}
		}
	}

}
