package practice

func areAlmostEqual(s1 string, s2 string) bool {

	n := len(s1)

	count := 0
	i, j := 0, 0

	for p := 0; p < n; p++ {

		if s1[p] != s2[p] {
			count++
			if count > 2 {
				return false
			} else {
				if count == 1 {
					i = p
				} else {
					j = p
				}
			}
		}
	}

	return s1[i] == s2[j] && s2[i] == s1[j]
}
