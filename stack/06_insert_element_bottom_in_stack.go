package stack

import "container/list"

func insertStackBottom(stack *list.List, data int) {

	if stack.Len() == 0 {
		stack.PushBack(data)
		return
	}

	pushBottom(stack, data)
}

func pushBottom(stack *list.List, data int) {

	// Base case
	if stack.Len() == 0 {
		stack.PushBack(data)
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	pushBottom(stack, data)

	stack.PushBack(back.Value)
}
