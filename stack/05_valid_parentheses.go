package stack

import "container/list"

func isValidParenthesis(str string) bool {

	stack := list.New()

	for i := 0; i < len(str); i++ {

		var ch byte = str[i]

		if ch == '{' || ch == '(' || ch == '[' {
			stack.PushBack(ch)
		} else {
			if stack.Len() > 0 {
				back := stack.Back()
				peek := back.Value.(byte)
				if (ch == '}' && peek == '{') ||
					(ch == ')' && peek == '(') ||
					(ch == ']' && peek == '[') {

					stack.Remove(back)
				} else {
					return false
				}
			} else {
				return false
			}

		}
	}

	return stack.Len() == 0
}
