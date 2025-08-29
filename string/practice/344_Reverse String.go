package practice

import "fmt"

func reverseByteArray(s []rune) {

	start := 0
	end := len(s) - 1

	var temp rune
	for start < end {
		temp = s[start]
		s[start] = s[end]
		s[end] = temp

		start++
		end--
	}

	fmt.Println(string(s))
}
