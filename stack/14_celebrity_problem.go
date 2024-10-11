package stack

import "container/list"

func celebrityProblem(m [][]int) int {

	size := len(m)
	stack := list.New()
	// insert all person in stack
	for i := 0; i < size; i++ {
		stack.PushBack(i)
	}

	var a int
	var b int

	for stack.Len() > 1 {
		// Pop top 2 person
		a = stack.Back().Value.(int)
		stack.Remove(stack.Back())
		b = stack.Back().Value.(int)
		stack.Remove(stack.Back())

		if m[a][b] == 1 {
			// a knows b, insert b again in stack
			stack.PushBack(b)
		} else {
			// b knows a, insert a again in stack
			stack.PushBack(a)
		}
	}

	// Single element in stack is potential celebrity
	ans := stack.Back().Value.(int)

	zeroCount := 0
	oneCount := 0

	for i := 0; i < size; i++ {
		if m[ans][i] == 0 {
			zeroCount++
		}
		if m[i][ans] == 1 {
			oneCount++
		}
	}

	if zeroCount != size {
		return -1
	}

	if oneCount != size-1 {
		return -1
	}

	return ans
}
