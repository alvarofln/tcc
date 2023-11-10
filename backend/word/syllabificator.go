package word

import (
	"gotoolapi/utils"
	"regexp"
	"strings"
)

var (
	digraphVowelPattern     = regexp.MustCompile(`^([bcdfgptvw][rl]|[lcntw]h|[gq]u)[aeiou]`) //detecta inicio da silaba que contem encontro consonantal inseparavel D + V
	consonantVowelPattern   = regexp.MustCompile(`^[^aeiou][aeiou]`)                         //detecta inicio da silaba que contem C + V
	nasalEOPattern          = regexp.MustCompile(`^(?:[ãõ]e|ão)`)                            // casos onde o e u são semivogais  V + S
	vowelIPattern           = regexp.MustCompile(`^i([nzlrm]([^aeiou]|$)|u)`)
	vowelUPattern           = regexp.MustCompile(`^u[nzlrm]([^aeiou]|$)`)
	hasVowelPattern         = regexp.MustCompile(`[aeiou]`)
	removeCharactersPattern = regexp.MustCompile(`[^a-záâãàéêíóôõúüç]`)
)

func Syllabificate(input string) []string {
	normalizedInput := removeCharactersPattern.ReplaceAllString(strings.ToLower(input), "")

	inputRunes := []rune(normalizedInput)
	inputUnaccentedRunes := []rune(utils.RemoveAccents(normalizedInput))

	hasVowel := false

	var tokens []Token
	lastToken := Token{}

	for i := 0; i < len(inputUnaccentedRunes); i++ {
		ch := string(inputRunes[i])                               //accented
		chUnaccented := string(inputUnaccentedRunes[i])           //unaccented
		currentWordRunes := inputRunes[i:]                        //accented
		currentWordUnaccented := string(inputUnaccentedRunes[i:]) //unaccented

		if digraphVowelPattern.MatchString(currentWordUnaccented) {
			if hasVowel {
				tokens = append(tokens, Token{Separator, "-"})
			}
			tokens = append(tokens, Token{Digraph, string(currentWordRunes[:2])})
			tokens = append(tokens, Token{Vowel, string(currentWordRunes[2:3])})
			hasVowel = true
			i += 2
			continue
		} else if consonantVowelPattern.MatchString(currentWordUnaccented) {
			if hasVowel {
				tokens = append(tokens, Token{Separator, "-"})
			}
			tokens = append(tokens, Token{Consonant, string(currentWordRunes[:1])})
			tokens = append(tokens, Token{Vowel, string(currentWordRunes[1:2])})
			hasVowel = true
			i += 1
			continue
		}

		if hasVowelPattern.MatchString(chUnaccented) {

			if len(tokens) > 0 {
				lastToken = tokens[len(tokens)-1]
			}

			if lastToken.Type == Vowel {
				if nasalEOPattern.MatchString(string(inputRunes[i-1:])) { // accented
					tokens = append(tokens, Token{SemiVowel, ch})
					continue
				}
				if !utils.HasAccent(ch) && (strings.HasPrefix(currentWordUnaccented, "i") && !vowelIPattern.MatchString(currentWordUnaccented) || strings.HasPrefix(currentWordUnaccented, "u") && !vowelUPattern.MatchString(currentWordUnaccented)) {
					tokens = append(tokens, Token{SemiVowel, ch})
					continue
				}
			}
			if lastToken.Type == SemiVowel || lastToken.Type == Vowel {
				tokens = append(tokens, Token{Separator, "-"})
			}

			tokens = append(tokens, Token{Vowel, ch})
			hasVowel = true
		} else {
			tokens = append(tokens, Token{Consonant, ch})
		}
	}

	var syllables []string
	syllable := ""

	for _, token := range tokens {
		if token.Type == Separator {
			syllables = append(syllables, syllable)
			syllable = ""
		} else {
			syllable += token.Value
		}
	}

	if len(syllable) > 0 {
		syllables = append(syllables, syllable)
	}

	return syllables
}
