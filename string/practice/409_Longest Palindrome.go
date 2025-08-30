package practice

func longestPalindromeOptimized(s string) int {
	set := make([]int, 128)

	oddCount := 0
	result := 0

	for _, ch := range s {
		if set[ch] != 0 {
			result += 2
			set[ch] = 0
			oddCount--
		} else {
			set[ch] = 1
			oddCount++
		}
	}

	if oddCount > 0 {
		result += 1
	}

	return result
}

func longestPalindrome(s string) int {
	freqMap := make(map[rune]int)

	oddCount := 0
	result := 0

	for _, ch := range s {
		freqMap[ch]++
		curFreq := freqMap[ch]
		if curFreq%2 == 0 {
			result += 2
			oddCount--
		} else {
			oddCount++
		}
	}

	if oddCount > 0 {
		result += 1
	}

	return result
}
