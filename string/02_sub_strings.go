package string

import (
	"fmt"
)

/**

	** Sub-Strings **

	Order and Continuty

	Ex: "abc"
	SubStrings:
	a
	ab
	abc
	b
	bc
	c

	Total: n(n+1)/2

**/

func printSubStrings(str string) {

	n := len(str)

	for start := 0; start < n; start++ {

		for end := start + 1; end <= n; end++ {

			fmt.Println(str[start:end])

		}
	}
}
