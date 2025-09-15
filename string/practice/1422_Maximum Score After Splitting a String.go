package practice

func maxScoreAfterSplitting(s string) int {

	n := len(s)

	oneCount := 0
	for i := range n {
		if s[i] == '1' {
			oneCount++
		}
	}

	ans := 0
	zeroCount := 0
	for i := range n - 1 {
		if s[i] == '0' {
			zeroCount++
		} else {
			oneCount--
		}

		ans = max(ans, oneCount+zeroCount)
	}

	return ans
}
