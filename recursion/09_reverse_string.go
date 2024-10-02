package recursion

func reverseString(str string) string {

	var runes []int32 = []rune(str)

	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		// swap values
		runes[start], runes[end] = runes[end], runes[start]
	}

	return string(runes)

}

func reverseStr(runes []rune, start, end int) {

	// Base case
	if start > end {
		return
	}

	// swap values
	runes[start], runes[end] = runes[end], runes[start]

	reverseStr(runes, start+1, end-1)
}
