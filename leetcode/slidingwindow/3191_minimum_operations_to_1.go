package slidingwindow

func minOperationsToEqual1(nums []int) int {

	size := len(nums)
	count := 0

	for i := 0; i < size-2; i++ {
		if nums[i] == 0 {
			nums[i] = 1 - nums[i]
			nums[i+1] = 1 - nums[i+1]
			nums[i+2] = 1 - nums[i+2]

			count++
		}
	}

	if nums[size-1] == 0 || nums[size-2] == 0 {
		return -1
	}

	return count
}
