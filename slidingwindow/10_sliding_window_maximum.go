package slidingwindow

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	result := []int{}

	if n == 0 {
		return result
	}

	deque := list.New()

	// first window
	index := 0
	for index < k {
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] <= nums[index] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(index)

		index++
	}

	if deque.Len() > 0 {
		result = append(result, nums[deque.Front().Value.(int)])
	}

	// other windows
	for i := 1; i < n-k+1; i++ {
		if deque.Len() > 0 && deque.Front().Value.(int) <= i-1 {
			deque.Remove(deque.Front())
		}

		for deque.Len() > 0 && nums[deque.Back().Value.(int)] <= nums[i+k-1] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(i + k - 1)

		result = append(result, nums[deque.Front().Value.(int)])
	}

	return result
}
