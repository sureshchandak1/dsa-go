package practice

func parseBoolExpr(expression string) bool {

	n := len(expression)
	stack := []rune{}

	for i := 0; i < n; i++ {
		ch := rune(expression[i])

		if ch == ')' {
			list := []rune{}

			for stack[len(stack)-1] != '(' {
				list = append(list, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // remove '('

			op := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans := evaluateExpression(list, op)
			stack = append(stack, ans)

		} else {
			if ch != ',' {
				stack = append(stack, ch)
			}
		}
	}

	return stack[len(stack)-1] == 't'
}

func evaluateExpression(list []rune, op rune) rune {
	if op == '&' {
		if findExpressionResult(list, 'f') {
			return 'f'
		}
		return 't'
	} else if op == '|' {
		if findExpressionResult(list, 't') {
			return 't'
		}
		return 'f'
	} else {
		// not operator
		if list[0] == 't' {
			return 'f'
		}
		return 't'
	}
}

func findExpressionResult(list []rune, val rune) bool {
	for _, ch := range list {
		if ch == val {
			return true
		}
	}
	return false
}
