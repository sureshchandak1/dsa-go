package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
}

func addStrings(num1 string, num2 string) string {

	var buffer bytes.Buffer

	carrey := 0

	for len(num1) != 0 || len(num2) != 0 || carrey != 0 {

		var n1 int = 0
		if len(num1) > 0 {
			ch1 := num1[len(num1)-1:]
			n1, _ = strconv.Atoi(ch1)
			num1 = num1[0 : len(num1)-1]
		}

		var n2 int = 0
		if len(num2) > 0 {
			ch2 := num2[len(num2)-1:]
			n2, _ = strconv.Atoi(ch2)
			num2 = num2[0 : len(num2)-1]
		}

		sum := n1 + n2 + carrey

		rem := strconv.Itoa(sum % 10)
		carrey = sum / 10

		buffer.WriteString(rem)

	}

	return reverseString(buffer.String())
}

func reverseString(str string) string {

	runes := []rune(str)

	for start, end := 0, len(str)-1; start < end; start, end = start+1, end-1 {
		runes[start], runes[end] = runes[end], runes[start]
	}

	return string(runes)

}
