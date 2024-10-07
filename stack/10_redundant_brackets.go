package stack

import (
	"container/list"
)

func findRedundantBrackets(str string) bool {

	stack := list.New()

	for i := 0; i < len(str); i++ {

		var ch byte = str[i]

		if ch == '(' || ch == '/' || ch == '*' || ch == '-' || ch == '+' {
			stack.PushBack(ch)
		} else {
			if ch == ')' {
				isRedundant := true
				for stack.Back().Value.(byte) != '(' {
					top := stack.Back()
					topVal := top.Value.(byte)
					if topVal == '/' || topVal == '*' || topVal == '-' || topVal == '+' {
						isRedundant = false
					}

					stack.Remove(top)
				}

				if isRedundant {
					return true
				}

				stack.Remove(stack.Back()) // remove opening brackets '('
			}
		}
	}

	return false
}
