package practice

func reversePrefix(word string, ch byte) string {

	runes := []rune(word)

	for i, currentCh := range runes {

		if currentCh == rune(ch) {
			reverseString(runes, 0, i)
			return string(runes)
		}
	}

	return word
}

func reverseString(runes []rune, start, end int) {

	for start < end {
		runes[start], runes[end] = runes[end], runes[start]

		start++
		end--
	}
}
