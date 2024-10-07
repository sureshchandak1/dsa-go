package stack

import "container/list"

func isValidParenthesis(str string) bool {

	runes := []rune(str)
	stack := list.New()

	for i := 0; i < len(runes); i++ {

		var ch int32 = runes[i]

		if ch == '{' || ch == '(' || ch == '[' {
			stack.PushBack(ch)
		} else {
			if stack.Len() > 0 {
				back := stack.Back()
				if (ch == '}' && back.Value == '{') ||
					(ch == ')' && back.Value == '(') ||
					(ch == ']' && back.Value == '[') {

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
