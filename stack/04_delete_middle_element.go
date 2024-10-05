package stack

import "container/list"

func deleteMiddleElement(stack *list.List) {

	size := stack.Len()

	mid := size / 2
	if size%2 == 0 {
		mid = mid - 1
	}

	deleteMiddle(stack, 0, mid)
}

func deleteMiddle(stack *list.List, count int, mid int) {

	// base case
	if count == mid {
		stack.Remove(stack.Back())
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	deleteMiddle(stack, count+1, mid)

	stack.PushBack(back.Value)

}
