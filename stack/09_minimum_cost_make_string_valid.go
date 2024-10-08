package stack

import "container/list"

func findMinimumCost(str string) int {

	size := len(str)

	if size%2 != 0 {
		return -1
	}

	stack := list.New()

	for i := 0; i < size; i++ {

		var ch byte = str[i]

		if ch == '{' {
			stack.PushBack(ch)
		} else {
			if stack.Len() > 0 && stack.Back().Value.(byte) == '{' {
				stack.Remove(stack.Back())
			} else {
				stack.PushBack(ch)
			}
		}
	}

	var a = 0 // open bracket count
	var b = 0 // close bracket count

	for stack.Len() != 0 {
		back := stack.Back()
		if back.Value.(byte) == '{' {
			a++
		} else {
			b++
		}

		stack.Remove(back)
	}

	result := ((a + 1) / 2) + ((b + 1) / 2)

	return result
}
