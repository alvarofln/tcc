package word

import (
	"regexp"
	"strings"
)

var (
	iuoPattern   = regexp.MustCompile(`.*([iu](s|m|ns)?|o(m|ns))$`)
	aeoPattern   = regexp.MustCompile(`.*([ae](s|m|ns)?|os?)$`)
	atonePattern = regexp.MustCompile(`^(o|a|e|os|ou|as|ao|do|da|de|em|no|nos|me|te|que|se|sem|com|vos)$`) //and many others :)
)

func DetectStress(syllables []string) StressType {
	if len(syllables) == 0 {
		return Atone
	}

	if len(syllables) == 1 {
		if atonePattern.MatchString(syllables[0]) {
			return Atone
		}
		return Oxytone
	}

	lastIndex := len(syllables) - 1

	for i := lastIndex; i >= lastIndex-3 && i >= 0; i-- {
		syllable := syllables[i]
		if strings.ContainsAny(syllable, "áâéêíóôú") {
			return StressType(lastIndex - i)
		}
	}

	lastSyllable := syllables[lastIndex]

	if strings.ContainsAny(lastSyllable, "ãõ") || iuoPattern.MatchString(lastSyllable) || !aeoPattern.MatchString(lastSyllable) {
		return Oxytone
	}

	if strings.Join(syllables, "") == "porque" { //exceção
		return Oxytone
	}

	return Paroxytone
}
