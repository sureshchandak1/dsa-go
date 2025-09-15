package practice

func vowelStrings(words []string, queries [][]int) []int {

	n := len(words)
	prefixSum := make([]int, n)

	prefixSum[0] = isVowelWord(words[0])
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + isVowelWord(words[i])
	}

	m := len(queries)
	ans := make([]int, m)

	for i := range m {

		l := queries[i][0]
		r := queries[i][1]

		res := prefixSum[r]
		if l != 0 {
			res -= prefixSum[l-1]
		}

		ans[i] = res
	}

	return ans
}

func isVowelWord(word string) int {

	first := rune(word[0])
	last := rune(word[len(word)-1])

	if (first == 'a' || first == 'e' || first == 'i' || first == 'o' || first == 'u') &&
		(last == 'a' || last == 'e' || last == 'i' || last == 'o' || last == 'u') {
		return 1
	}

	return 0
}
