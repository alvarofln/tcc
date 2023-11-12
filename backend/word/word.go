package word

type StressType int

const (
	Atone StressType = iota - 1
	Oxytone
	Paroxytone
	Proparoxytone
)

func (sc StressType) String() string {
	switch sc {
	case Atone:
		return "Atone"
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
