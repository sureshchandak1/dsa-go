package practice

import "strings"

func reverseParenthesesStringOptimized(s string) string {

	n := len(s)
	stack := make([]int, 0)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		ch := s[i]
		if ch == '(' {
			stack = append(stack, i)
		} else if ch == ')' {
			openIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr[openIndex] = i
			arr[i] = openIndex
		}
	}

	var result strings.Builder
	direction := 1
	i := 0
	for i < n {
		ch := s[i]

		if ch == '(' || ch == ')' {
			i = arr[i]
			direction = -direction
		} else {
			result.WriteByte(ch)
		}

		i += direction
	}

	return result.String()
}

func reverseParenthesesString(s string) string {

	stack := make([]rune, 0)
	var result string

	for _, ch := range s {
		if ch == '(' || ch != ')' {
			stack = append(stack, ch)
		} else {
			list := make([]rune, 0)

			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				list = append(list, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// remove the opening bracket
			stack = stack[:len(stack)-1]

			stack = append(stack, list...)
		}
	}

	for len(stack) > 0 {
		result = string(stack[len(stack)-1]) + result
		stack = stack[:len(stack)-1]
	}

	return result
}
