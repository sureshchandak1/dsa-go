package queue

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {

	ans := []int{}

	size := len(nums)
	deque := list.New()

	i := 0 // start
	j := 0 // end

	for j < size {

		for deque.Len() > 0 && deque.Back().Value.(int) < nums[j] {
			deque.Remove(deque.Back())
		}
		deque.PushBack(nums[j])

		if j-i+1 < k {
			j++
		} else if j-i+1 == k {
			front := deque.Front()
			frontVal := front.Value.(int)

			ans = append(ans, frontVal)

			if nums[i] == frontVal {
				deque.Remove(front)
			}

			i++
			j++
		}
	}

	return ans
}
