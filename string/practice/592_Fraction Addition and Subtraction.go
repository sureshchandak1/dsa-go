package practice

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func fractionAddition(expression string) string {

	n := len(expression)

	num := 0
	den := 1

	i := 0
	for i < n {

		curNum := 0
		curDen := 0
		isNeg := false

		ch := rune(expression[i])

		if ch == '+' || ch == '-' {
			if ch == '-' {
				isNeg = true
			}
			i++
		}

		// form the num
		start := i
		for i < n && unicode.IsDigit(rune(expression[i])) {
			i++
		}
		curNum, _ = strconv.Atoi(expression[start:i])
		if isNeg {
			curNum *= -1
		}

		i++ // Skip '/' char

		// form the den
		start = i
		for i < n && unicode.IsDigit(rune(expression[i])) {
			i++
		}
		curDen, _ = strconv.Atoi(expression[start:i])

		num = num*curDen + curNum*den
		den = den * curDen
	}

	gcd := int(math.Abs(float64(getGCD(num, den))))

	num /= gcd
	den /= gcd

	return fmt.Sprintf("%d/%d", num, den)

}

func getGCD(a, b int) int {

	// Base case
	if a == 0 {
		return b
	}

	return getGCD(b%a, a)
}
