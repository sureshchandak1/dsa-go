package practice

import "math"

func makeGoodString(s string) string {

	var sb []rune

	for _, ch := range s {
		len := len(sb)

		if len > 0 && int(math.Abs(float64(sb[len-1]-ch))) == 32 {
			sb = sb[:len-1]
		} else {
			sb = append(sb, ch)
		}
	}

	return string(sb)
}

func makeGoodStringStack(s string) string {

	stack := []rune{}

	for _, ch := range s {
		if len(stack) > 0 && int(math.Abs(float64(stack[len(stack)-1]-ch))) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	ans := ""
	for len(stack) > 0 {
		ans = string(stack[len(stack)-1]) + ans
		stack = stack[:len(stack)-1]
	}

	return ans
}
