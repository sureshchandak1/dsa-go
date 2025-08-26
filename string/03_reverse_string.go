package string

func reverseString(str string) string {

	runes := []rune(str)

	start := 0
	end := len(runes) - 1

	for start < end {

		runes[start], runes[end] = runes[end], runes[start]

		start++
		end--
	}

	return string(runes)
}

func reverseString1(str string) string {

	runes := []rune(str)
	n := len(runes)

	for i := range n / 2 {

		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}
