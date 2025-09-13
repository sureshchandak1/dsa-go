package practice

import "strings"

func findLongestSpecialSubstring(s string) int {

	n := len(s)
	mapCount := make(map[string]int)

	for i := 0; i < n; i++ {
		var curSb strings.Builder
		lastChar := s[i]

		for j := i; j < n; j++ {
			if curSb.Len() == 0 || s[j] == lastChar {
				curSb.WriteByte(s[j])
				mapCount[curSb.String()]++
			} else {
				break
			}
		}
	}

	maxLen := -1

	for str, freq := range mapCount {
		if freq > 2 {
			if len(str) > maxLen {
				maxLen = len(str)
			}
		}
	}

	return maxLen
}
