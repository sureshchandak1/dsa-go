package practice

func stringMatching(words []string) []string {

	n := len(words)
	result := []string{}

	for i := range n {
		for j := 0; j < n; j++ {
			if i == j || len(words[i]) > len(words[j]) {
				continue
			}

			if isSubstring(words[i], words[j]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}

func isSubstring(sub, word string) bool {
	n := len(word)
	m := len(sub)

	for j := 0; j < n; j++ {
		isMatch := true
		k := j
		i := 0
		for i < m {
			if k >= len(word) || word[k] != sub[i] {
				isMatch = false
				break
			}

			k++
			i++
		}

		if isMatch {
			return true
		}
	}

	return false
}
