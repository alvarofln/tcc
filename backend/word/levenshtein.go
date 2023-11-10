package word

// Adapted from https://rosettacode.org/wiki/Levenshtein_distance#Go
func Levenshtein(s, t string) int {
	runesS := []rune(s)
	runesT := []rune(t)
	lenS := len(runesS)
	lenT := len(runesT)

	d := make([][]int, lenS+1)
	for i := range d {
		d[i] = make([]int, lenT+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= lenT; j++ {
		for i := 1; i <= lenS; i++ {
			if runesS[i-1] == runesT[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
		}

	}
	return d[lenS][lenT]
}

func min(values ...int) int {
	minVal := values[0]
	for _, v := range values {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}
