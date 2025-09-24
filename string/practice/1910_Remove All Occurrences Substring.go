package practice

import "strings"

func removeOccurrences(s string, part string) string {

	n := len(s)
	m := len(part)
	var sb strings.Builder

	for i := range n {

		sb.WriteByte(s[i])

		if sb.Len() >= m {
			sub := sb.String()[sb.Len()-m:]
			if sub == part {
				newStr := sb.String()[:sb.Len()-m]

				sb.Reset()
				sb.WriteString(newStr)
			}
		}
	}

	return sb.String()
}
