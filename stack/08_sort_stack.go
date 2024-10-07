package stack

import "container/list"

func sortStack(stack *list.List) {

	// Base case
	if stack.Len() == 0 {
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	sortStack(stack)

	sortedInsert(stack, back.Value.(int))
}

func sortedInsert(stack *list.List, data int) {

	// Base case
	if stack.Len() == 0 || (stack.Back().Value.(int) < data) {
		stack.PushBack(data)
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	sortedInsert(stack, data)

	stack.PushBack(back.Value)
}
