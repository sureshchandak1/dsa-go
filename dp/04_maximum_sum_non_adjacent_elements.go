package dp

import "slices"

func maxSumNonAdjacentElements(nums []int) int {
	// n := len(nums)
	// return solveMaxSumNonAdjacentElements(nums, n-1)

	// n := len(nums)
	// dp := slices.Repeat([]int{-1}, n+1)
	// return solveMaxSumNonAdjacentElementsDP(nums, n-1, dp)

	// return solveMaxSumNonAdjacentElementsTab(nums)

	return solveMaxSumNonAdjacentElementsOp(nums)
}

// Recursion
func solveMaxSumNonAdjacentElements(nums []int, n int) int {

	// Base case
	if n < 0 {
		return 0
	}
	if n == 0 {
		return nums[0]
	}

	incl := solveMaxSumNonAdjacentElements(nums, n-2) + nums[n]

	excl := solveMaxSumNonAdjacentElements(nums, n-1)

	return max(incl, excl)

}

// DP
func solveMaxSumNonAdjacentElementsDP(nums []int, n int, dp []int) int {

	// Base case
	if n < 0 {
		return 0
	}
	if n == 0 {
		return nums[0]
	}

	if dp[n] != -1 {
		return dp[n]
	}

	incl := solveMaxSumNonAdjacentElementsDP(nums, n-2, dp) + nums[n]

	excl := solveMaxSumNonAdjacentElementsDP(nums, n-1, dp)

	dp[n] = max(incl, excl)

	return dp[n]

}

// Tabulation
func solveMaxSumNonAdjacentElementsTab(nums []int) int {

	n := len(nums)
	dp := slices.Repeat([]int{0}, n)

	// according base case
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		incl, excl := 0, 0

		if i-2 >= 0 {
			incl = dp[i-2] + nums[i]
		}

		excl = dp[i-1]

		dp[i] = max(incl, excl)
	}

	return dp[n-1]

}

// Optimized
func solveMaxSumNonAdjacentElementsOp(nums []int) int {

	n := len(nums)

	prev2 := 0
	prev1 := nums[0]

	for i := 1; i < n; i++ {
		incl, excl := 0, 0

		incl = prev2 + nums[i]

		excl = prev1

		ans := max(incl, excl)

		prev2 = prev1
		prev1 = ans
	}

	return prev1

}
