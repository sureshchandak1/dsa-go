package stack

import "container/list"

func nextSmallerElement(arr []int) []int {

	stack := list.New()
	stack.PushBack(-1)

	ans := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {

		curr := arr[i]

		for stack.Back().Value.(int) >= curr {
			stack.Remove(stack.Back())
		}

		ans[i] = stack.Back().Value.(int)

		stack.PushBack(curr)
	}

	return ans
}