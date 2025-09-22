package practice

func minimumLengthAfterOperations(s string) int {

	n := len(s)
	freq := make([]int, 26)

	for i := 0; i < n; i++ {
		freq[s[i]-'a']++
	}

	count := 0
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			if freq[i]%2 == 0 {
				count += 2
			} else {
				count += 1
			}
		}
	}

	return count
}
