package queue

import "container/list"

func firstNegativeIntegerInWindow(arr []int, windowSize int) []int {

	ans := []int{}

	deque := list.New()

	for index := 0; index < windowSize; index++ {
		if arr[index] < 0 {
			deque.PushBack(index)
		}
	}

	// store first window result
	if deque.Len() > 0 {
		ans = append(ans, arr[deque.Front().Value.(int)])
	} else {
		ans = append(ans, 0)
	}

	for index := windowSize; index < len(arr); index++ {

		// Remove
		if deque.Len() > 0 && index-deque.Front().Value.(int) >= windowSize {
			deque.Remove(deque.Front())
		}

		// addition
		if arr[index] < 0 {
			deque.PushBack(index)
		}

		// store ans
		if deque.Len() > 0 {
			ans = append(ans, arr[deque.Front().Value.(int)])
		} else {
			ans = append(ans, 0)
		}
	}

	return ans
}
