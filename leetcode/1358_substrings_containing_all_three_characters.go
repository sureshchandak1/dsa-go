package leetcode

func numberOfSubstrings(s string) int {
	size := len(s)
	var freq [3]int
	count := 0
	curCount := 0

	var ch byte
	left := 0
	for right := range size {
		// expantion
		ch = s[right]
		freq[ch-'a']++
		if freq[ch-'a'] == 1 {
			curCount++
		}

		// shrking
		for curCount == 3 {
			count += size - right
			ch = s[left]
			freq[ch-'a']--
			if freq[ch-'a'] == 0 {
				curCount--
			}

			left++
		}
	}

	return count
}
