package practice

// Sliding window solution
func permutationInStringOptimized(s1 string, s2 string) bool {

	n := len(s1)
	m := len(s2)

	if m < n {
		return false
	}

	map1 := make([]int, 26)
	map2 := make([]int, 26)

	for i := 0; i < n; i++ {
		map1[s1[i]-'a']++
		map2[s2[i]-'a']++
	}

	// init window
	if isMatched(map1, map2) {
		return true
	}

	for i := 1; i <= m-n; i++ {
		map2[s2[i-1]-'a']--   // deletion
		map2[s2[i+n-1]-'a']++ // addition

		if isMatched(map1, map2) {
			return true
		}
	}

	return false

}

func permutationInString(s1 string, s2 string) bool {

	n := len(s1)
	m := len(s2)

	if m < n {
		return false
	}

	map1 := make([]int, 26)

	for i := 0; i < n; i++ {
		map1[s1[i]-'a']++
	}

	for i := 0; i <= m-n; i++ {
		map2 := make([]int, 26)

		for j := 0; j < n; j++ {
			map2[s2[i+j]-'a']++
		}

		if isMatched(map1, map2) {
			return true
		}
	}

	return false
}

func isMatched(map1 []int, map2 []int) bool {
	for i := 0; i < 26; i++ {
		if map1[i] != map2[i] {
			return false
		}
	}

	return true
}
