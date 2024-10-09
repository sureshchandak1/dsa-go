package stack

import "container/list"

func largestRectangle(heights []int) int {

	size := len(heights)

	next := nextSmallerElementIndexs(heights)
	prev := prevSmallerElementIndexs(heights)

	area := -1

	for i := 0; i < size; i++ {
		l := heights[i]

		if next[i] == -1 {
			next[i] = size
		}

		b := next[i] - prev[i] - 1

		newArea := l * b

		area = max(area, newArea)
	}

	return area
}

func nextSmallerElementIndexs(heights []int) []int {

	stack := list.New()
	stack.PushBack(-1)

	ans := make([]int, len(heights))

	for i := len(heights) - 1; i >= 0; i-- {

		curr := heights[i]

		for stack.Back().Value.(int) != -1 && heights[stack.Back().Value.(int)] >= curr {
			stack.Remove(stack.Back())
		}

		ans[i] = stack.Back().Value.(int)

		stack.PushBack(i)
	}

	return ans
}

func prevSmallerElementIndexs(heights []int) []int {

	stack := list.New()
	stack.PushBack(-1)

	ans := make([]int, len(heights))

	for i := 0; i < len(heights); i++ {

		curr := heights[i]

		for stack.Back().Value.(int) != -1 && heights[stack.Back().Value.(int)] >= curr {
			stack.Remove(stack.Back())
		}

		ans[i] = stack.Back().Value.(int)

		stack.PushBack(i)
	}

	return ans
}
