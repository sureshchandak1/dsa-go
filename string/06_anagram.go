package string

func isAnagram(a string, b string) bool {

	if len(a) != len(b) {
		return false
	}

	freq := make([]byte, 26)

	for i := range len(a) {
		freq[a[i]-'a']++
		freq[b[i]-'a']--
		// If same frequency then count is 0
	}

	for i := range 26 {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}
