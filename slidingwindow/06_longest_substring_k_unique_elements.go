package slidingwindow

func longestKUniqueSubStr(str string, k int) int {
	n := len(str)
	maxLen := -7687687643

	windowStart := 0
	windowEnd := 0
	len := 0
	freqHash := make([]int, 26)
	uniqueCharCount := 0

	var index byte
	for windowEnd < n {
		// Expansion face
		index = str[windowEnd] - 'a' // a = 97

		if freqHash[index] == 0 {
			uniqueCharCount++
			freqHash[index]++
		} else {
			freqHash[index]++
		}

		if uniqueCharCount == k {
			len = windowEnd - windowStart + 1
			maxLen = max(maxLen, len)
		} else if uniqueCharCount > k {
			// Shrinking face
			for windowStart < windowEnd && uniqueCharCount > k {

				index = str[windowStart] - 'a'

				freqHash[index]--
				if freqHash[index] == 0 {
					uniqueCharCount--
				}

				windowStart++
			}
		}

		windowEnd++
	}

	if maxLen == -7687687643 {
		return -1
	} else {
		return maxLen
	}

}
