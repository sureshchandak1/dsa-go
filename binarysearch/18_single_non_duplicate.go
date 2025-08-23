package binarysearch

func singleNonDuplicate(nums []int) int {

	start := 0
	end := len(nums) - 2

	for start <= end {

		mid := start + (end-start)/2

		if nums[mid] == nums[mid^1] {
			// Left half -> move towards right
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return nums[start]
}
