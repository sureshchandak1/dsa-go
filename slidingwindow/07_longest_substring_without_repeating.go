package slidingwindow

func lengthOfLongestSubstringOptimized(s string) int {
	size := len(s)
	windowStart, windowEnd := 0, 0
	indexMap := make(map[byte]int)

	maxLen := -23423423423

	var ch byte
	for windowEnd < size {
		ch = s[windowEnd]

		index, exits := indexMap[ch]
		if exits && index >= windowStart {
			// Shrikinig
			windowStart = index + 1

		}

		indexMap[ch] = windowEnd
		maxLen = max(maxLen, windowEnd-windowStart+1)

		windowEnd++
	}

	if maxLen == -23423423423 {
		return 0
	}
	return maxLen
}

func lengthOfLongestSubstring1(s string) int {
	size := len(s)
	windowStart, windowEnd := 0, 0
	set := make(map[byte]bool)

	maxLen := -23423423423

	var ch byte
	for windowEnd < size {
		ch = s[windowEnd]

		_, exits := set[ch]
		if exits {
			// Shrikinig
			for _, ok := set[ch]; windowStart < windowEnd && ok; {
				delete(set, s[windowStart])
				windowStart++
				_, ok = set[ch]
			}

		}

		set[ch] = true
		maxLen = max(maxLen, windowEnd-windowStart+1)

		windowEnd++
	}

	if maxLen == -23423423423 {
		return 0
	}
	return maxLen
}
