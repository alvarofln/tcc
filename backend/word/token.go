package word

type TokenType int

const (
	Vowel TokenType = iota + 1
	SemiVowel
	Consonant
	Separator
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case Vowel:
		return "[V]"
	case SemiVowel:
		return "[S]"
	case Consonant:
		return "[C]"
	case Separator:
		return "[SEP]"
	default:
		return "[?]"
	}
}
