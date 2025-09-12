package practice

import "strings"

func addSpaces(s string, spaces []int) string {

	n := len(s)
	m := len(spaces)
	sb := strings.Builder{}
	sb.Grow(n + m)

	i, j := 0, 0
	for i < n {
		if j < m && i == spaces[j] {
			sb.WriteByte(' ')
			j++
		}
		sb.WriteByte(s[i])
		i++
	}

	return sb.String()
}
