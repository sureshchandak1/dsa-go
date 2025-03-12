package slidingwindow

import "container/list"

func firstNegativeNumberK(nums []int, k int) []int {
	return firstNegativeNumberKOptimized(nums, k)
}

func firstNegativeNumberKOptimized(nums []int, k int) []int {

	n := len(nums)
	result := []int{}

	queue := list.New()

	// first  window
	index := 0
	for index < k {
		if nums[index] < 0 {
			queue.PushBack(nums[index])
		}
		index++
	}

	if queue.Len() != 0 {
		result = append(result, queue.Front().Value.(int))
	}

	// next window
	for index = 1; index < n-k+1; index++ {
		// Remove
		if nums[index-1] < 0 {
			queue.Remove(queue.Front())
		}

		// add
		if nums[index+k-1] < 0 {
			queue.PushBack(nums[index+k-1])
		}

		if queue.Len() != 0 {
			result = append(result, queue.Front().Value.(int))
		}
	}

	return result
}

func firstNegativeNumberKBrutForce(nums []int, k int) []int {
	n := len(nums)

	result := []int{}

	for i := 0; i < n-k+1; i++ {
		for j := 0; j < k; j++ {
			if nums[i+j] < 0 {
				result = append(result, nums[i+j])
			}
		}
	}

	return result
}
