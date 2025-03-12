package slidingwindow

/**
Leetcode: 2461. Maximum Sum of Distinct Subarrays With Length K
*/
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	countMap := make(map[int]int)

	sum := 0
	maxSum := 0

	index := 0
	// first window
	for index < n && index < k {
		sum += nums[index]
		countMap[nums[index]]++
		index++
	}

	if len(countMap) == k {
		maxSum = sum
	}

	// next window
	for index := 1; index < n-k+1; index++ {

		countMap[nums[index-1]]--
		if countMap[nums[index-1]] == 0 {
			delete(countMap, nums[index-1])
		}

		sum = sum - nums[index-1] + nums[index+k-1]
		countMap[nums[index+k-1]]++

		if len(countMap) == k {
			maxSum = max(maxSum, sum)
		}
	}

	return int64(maxSum)
}
