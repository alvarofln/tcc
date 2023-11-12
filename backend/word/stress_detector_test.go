package word

import (
	"testing"
)

func TestOxytones(t *testing.T) {
	oxytones := []string{"andar", "viver", "atum", "sorrir", "computador", "magoar", "ruir", "anzol", "papel", "papéis", "cantil", "cantis", "bombom", "bom", "amou", "chapéu", "viveu", "país", "paz", "pais", "sabiá", "dó", "legal", "cristal", "ruir", "abacaxi", "também", "alguém", "Abigail", "algum", "algoz", "raiz", "filé", "metrô", "urubu", "ali", "purê"}

	for _, oxytone := range oxytones {
		syllables := Syllabificate(oxytone)

		if DetectStress(syllables) != Oxytone {
			t.Errorf("DetectStress(%v) = %v; expected %v", syllables, DetectStress(syllables), Oxytone)
		}
	}
}
func TestParoxytones(t *testing.T) {
	paroxytones := []string{"anda", "vive", "sorria", "computa", "amargo", "magoa", "corroe", "marca", "casa", "anjo", "canja", "geleia", "tênis", "voo", "enjoo", "hífen", "alcateia", "joia", "costa", "sótão", "incrível", "boa", "voa", "voam", "voamos", "rubrica", "têxtil", "teste", "pegada", "gado", "caldo", "Valdo", "alto", "câncer", "conga", "queijo", "doce", "dócil", "dóceis", "samba", "enredo", "pagode", "heroico", "inconstitucionalissimamente"}

	for _, paroxytone := range paroxytones {
		syllables := Syllabificate(paroxytone)

		if DetectStress(syllables) != Paroxytone {
			t.Errorf("DetectStress(%v) = %v; expected %v", syllables, DetectStress(syllables), Paroxytone)
		}
	}
}

func TestProparoxytones(t *testing.T) {
	proparoxytones := []string{"abóbora", "ângulo", "arquétipo", "árvore", "átomo", "Álvaro", "Bárbara", "básico", "bígamo", "brócolis", "bússola", "científico", "cítara", "círculo", "cômico", "crítico", "décima", "didático", "dinâmica", "dízimo", "dúvida", "época", "exército", "fábula", "fanático", "ginástica", "gótico", "harmônica", "hóspede", "índice", "ínterim", "jornalístico", "Júpiter", "kartódromo", "lâmina", "proparoxítona"}
	for _, proparoxytone := range proparoxytones {
		syllables := Syllabificate(proparoxytone)

		if DetectStress(syllables) != Proparoxytone {
			t.Errorf("DetectStress(%v) = %v; expected %v", syllables, DetectStress(syllables), Proparoxytone)
		}
	}
}

func TestAtones(t *testing.T) {
	atones := []string{"a", "e", "o", "as", "os", "ao"}
	for _, atone := range atones {
		syllables := Syllabificate(atone)

		if DetectStress(syllables) != Atone {
			t.Errorf("DetectStress(%v) = %v; expected %v", syllables, DetectStress(syllables), Atone)
		}
	}
}
