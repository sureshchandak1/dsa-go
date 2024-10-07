package stack

import "container/list"

func reverseStack(stack *list.List) {

	// Base case
	if stack.Len() == 0 {
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive call
	reverseStack(stack)

	insertAtBottom(stack, back.Value.(int))
}

func insertAtBottom(stack *list.List, data int) {

	// Base case
	if stack.Len() == 0 {
		stack.PushBack(data)
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	insertAtBottom(stack, data)

	stack.PushBack(back.Value)
}
