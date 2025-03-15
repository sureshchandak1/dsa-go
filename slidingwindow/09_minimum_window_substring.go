package slidingwindow

func minimumWindowSubstring(s string, t string) string {

	freqMap := make(map[byte]int)
	// Populate the map with t string
	for i := range len(t) {
		freqMap[t[i]]++
	}

	uniqueCharCount := len(freqMap)
	startIndex := -1
	windowStart, windowEnd := 0, 0
	minLen := 2523423432

	n := len(s)

	var ch byte
	var len int
	for windowEnd < n {
		// Expansion Phase
		ch = s[windowEnd]
		_, ok := freqMap[ch]
		if ok {
			freqMap[ch]--

			if freqMap[ch] == 0 {
				uniqueCharCount--
			}
		}

		// Shrinking Phase
		for uniqueCharCount == 0 {
			// find len
			len = windowEnd - windowStart + 1
			if len < minLen {
				minLen = len
				startIndex = windowStart
			}

			ch = s[windowStart]
			_, ok := freqMap[ch]
			if ok {
				freqMap[ch]++

				if freqMap[ch] > 0 {
					uniqueCharCount++
				}
			}

			windowStart++
		}

		windowEnd++
	}

	if startIndex == -1 {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
