package practice

func appendCharacters(s string, t string) int {

	n := len(s)
	m := len(t)

	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}
	}

	return m - j
}
