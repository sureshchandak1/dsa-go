package string

func reverseWords3(str string) string {

	runes := []rune(str)

	n := len(runes)

	start := 0

	for end := 0; end <= n; end++ {

		if end == n || runes[end] == ' ' {
			reverseWithStartEnd(runes, start, end-1)
			start = end + 1
		}
	}

	return string(runes)

}

func reverseWithStartEnd(runes []rune, start, end int) {

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]

		start++
		end--
	}
}
