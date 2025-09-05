package practice

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {

	count := make(map[string]int)

	words1 := strings.Fields(s1)
	words2 := strings.Fields(s2)

	for _, word := range words1 {
		count[word]++
	}

	for _, word := range words2 {
		count[word]++
	}

	var result []string
	for word, freq := range count {
		if freq == 1 {
			result = append(result, word)
		}
	}

	return result
}
