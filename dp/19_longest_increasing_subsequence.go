package dp

import "slices"

func lengthOfLIS(nums []int) int {
	// return lengthOfLISRes(nums, 0, -1)

	// n := len(nums)
	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, n+1)
	// }

	// return lengthOfLISDp(nums, 0, -1, dp)

	// return lengthOfLISTab(nums)

	// return lengthOfLISSC(nums)

	return lengthOfLISBinarySearch(nums)
}

// Recursion
func lengthOfLISRes(nums []int, curr, prev int) int {

	// base case
	if curr == len(nums) {
		return 0
	}

	include := 0
	if prev == -1 || nums[curr] > nums[prev] {
		include = 1 + lengthOfLISRes(nums, curr+1, curr)
	}

	exclude := lengthOfLISRes(nums, curr+1, prev)

	return max(include, exclude)
}

// Recursion + Memorization
// T.C = O(n * n), S.C = O(n * n)
func lengthOfLISDp(nums []int, curr, prev int, dp [][]int) int {

	// base case
	if curr == len(nums) {
		return 0
	}

	if dp[curr][prev+1] != -1 {
		return dp[curr][prev]
	}

	include := 0
	if prev == -1 || nums[curr] > nums[prev] {
		include = 1 + lengthOfLISDp(nums, curr+1, curr, dp)
	}

	exclude := lengthOfLISDp(nums, curr+1, prev, dp)

	dp[curr][prev+1] = max(include, exclude)

	return dp[curr][prev+1]
}

// Tabulation: Bottom-Up
// T.C = O(n * n), S.C = O(n * n)
func lengthOfLISTab(nums []int) int {

	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, n+1)
	}

	for curr := n - 1; curr >= 0; curr-- {

		for prev := curr - 1; prev >= -1; prev-- {

			include := 0
			if prev == -1 || nums[curr] > nums[prev] {
				include = 1 + dp[curr+1][curr+1]
			}

			exclude := dp[curr+1][prev+1]

			dp[curr][prev+1] = max(include, exclude)

		}
	}

	return dp[0][0]

}

// Space Optimization
// T.C = O(n * n), S.C = O(n)
func lengthOfLISSC(nums []int) int {

	n := len(nums)

	currRow := make([]int, n+1)
	nextRow := make([]int, n+1)

	for curr := n - 1; curr >= 0; curr-- {

		for prev := curr - 1; prev >= -1; prev-- {

			include := 0
			if prev == -1 || nums[curr] > nums[prev] {
				include = 1 + nextRow[curr+1]
			}

			exclude := nextRow[prev+1]

			currRow[prev+1] = max(include, exclude)

		}

		nextRow = slices.Clone(currRow)
	}

	return nextRow[0]

}

// DP + Binary Search
// T.C = O(n log n)
func lengthOfLISBinarySearch(nums []int) int {

	n := len(nums)

	if n == 0 {
		return 0
	}

	ans := []int{}
	ans = append(ans, nums[0])

	for index := 1; index < n; index++ {

		if nums[index] > ans[len(ans)-1] {
			ans = append(ans, nums[index])
		} else {
			// find index of just bigger element in ans
			rIndex := lowerBound(ans, nums[index])
			ans[rIndex] = nums[index]
		}
	}

	return len(ans)

}

func lowerBound(arr []int, element int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start < end {

		if arr[mid] < element {
			start = mid + 1
		} else {
			end = mid
		}

		mid = start + (end-start)/2
	}

	return start
}
