package practice

import "math"

func scoreOfString(s string) int {

	score := 0

	for i := range len(s) - 1 {
		score += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}

	return score
}
