package practice

import "strings"

func countPrefixSuffixPairs(words []string) int {

	n := len(words)
	count := 0

	for i := range n - 1 {

		str1 := words[i]
		for j := i + 1; j < n; j++ {
			str2 := words[j]
			if len(str1) > len(str2) {
				continue
			}

			if strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1) {
				count++
			}
		}
	}

	return count
}
