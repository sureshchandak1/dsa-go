package practice

func commonChars(words []string) []string {

	n := len(words)

	minFreq := make([]int, 26)
	firstWord := words[0]
	for _, ch := range firstWord {
		index := ch - 'a'
		minFreq[index]++
	}

	for i := 1; i < n; i++ {
		freq := make([]int, 26)
		for _, ch := range words[i] {
			freq[ch-'a']++
		}

		// update min freq
		for j := range 26 {
			minFreq[j] = int(min(float64(minFreq[j]), float64(freq[j])))
		}
	}

	result := []string{}

	for i := range 26 {
		if minFreq[i] != 0 {

			ch := rune(i + 'a')
			count := minFreq[i]
			for count > 0 {
				result = append(result, string(ch))
				count--
			}

		}
	}

	return result
}
