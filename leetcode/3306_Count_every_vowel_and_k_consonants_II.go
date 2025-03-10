package leetcode

func countOfSubstrings(word string, k int) int64 {
	return atLeast(word, k) - atLeast(word, k+1)
}

func atLeast(word string, k int) int64 {
	var count int = 0
	curConsonant := 0
	var freq [26]int
	size := len(word)

	left := 0
	for right := 0; right < size; right++ {
		ch := word[right]
		if isConsonant(ch) {
			curConsonant++
		}
		freq[ch-'a']++

		for curConsonant >= k && isAllVowelPresent(freq) {
			count += (size - right)
			c := word[left]
			if isConsonant(c) {
				curConsonant--
			}

			freq[c-'a']--

			left++
		}
	}

	return int64(count)
}

func isConsonant(ch byte) bool {
	return (ch != 'a' && ch != 'e' && ch != 'i' && ch != 'o' && ch != 'u')
}

func isAllVowelPresent(freq [26]int) bool {
	return (freq['a'-'a'] > 0 && freq['e'-'a'] > 0 && freq['i'-'a'] > 0 && freq['o'-'a'] > 0 && freq['u'-'a'] > 0)
}
