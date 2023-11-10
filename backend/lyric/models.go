package lyric

import "gotoolapi/word"

type InputLyric struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type LyricLine struct {
	Words         []*word.Word `json:"words"`
	SyllableCount int          `json:"syllableCount"`
}

type Lyric struct {
	Lines []*LyricLine `json:"lines"`
}
