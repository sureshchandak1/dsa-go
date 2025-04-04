package hashmap

import "sort"

func longestConsecutive2(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	set := make(map[int]bool)

	for _, val := range nums {
		set[val] = true
	}

	maxLen := 0

	for key := range set {
		prevEl := key - 1

		if !set[prevEl] {
			len := 1
			nextEl := key + 1

			for set[nextEl] {
				len++
				nextEl++
			}

			maxLen = max(maxLen, len)
		}
	}

	return maxLen
}

// n + log n
func longestConsecutive1(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxLen := 1
	prev := nums[0]
	currLen := 1

	for index := 1; index < len(nums); index++ {

		if prev == nums[index] {
			continue
		}

		if nums[index] == (prev + 1) {
			currLen++
			prev = nums[index]
		} else {
			currLen = 1
			prev = nums[index]
		}

		maxLen = max(maxLen, currLen)
	}

	return maxLen
}
