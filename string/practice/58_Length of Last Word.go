package practice

/**
func lengthOfLastWord(s string) int {

	n := len(s)
	start := n - 1

	for start >= 0 {

		for start >= 0 && s[start] == ' ' {
			start--
		}

		endIndex := start

		for start >= 0 && s[start] != ' ' {
			start--
		}

		return endIndex - start
	}

	return -1
}
**/

func lengthOfLastWord(s string) int {

	lastWorldLen := 0
	i := len(s) - 1

	for i >= 0 {
		if s[i] == ' ' {
			if lastWorldLen == 0 {
				i--
			} else {
				break
			}
		} else {
			lastWorldLen++
			i--
		}
	}

	return lastWorldLen
}
