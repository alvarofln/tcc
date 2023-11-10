package word

type StressType int

const (
	Oxytone       StressType = iota
	Paroxytone               = 1
	Proparoxytone            = 2
)

func (sc StressType) String() string {
	switch sc {
	case Oxytone:
		return "Oxytone"
	case Paroxytone:
		return "Paroxytone"
	case Proparoxytone:
		return "Proparoxytone"
	default:
		return "Unknown"
	}
}

type Word struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Syllables  []string   `json:"syllables"`
	StressType StressType `json:"stressType"`
}

func NewWord(id int64, name string, syllables []string, stressType StressType) *Word {
	return &Word{
		ID:         id,
		Name:       name,
		Syllables:  syllables,
		StressType: stressType,
	}
}

type SimilarWord struct {
	Word
	Similarity float64 `json:"similarity"`
}
