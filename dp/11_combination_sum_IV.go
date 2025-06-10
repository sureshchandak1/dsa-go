package dp

import "slices"

func targetSumWays(num []int, target int) int {
	// return targetSumWaysRes(num, target)

	// dp := slices.Repeat([]int{-1}, target+1)
	// return targetSumWaysDp(num, target, dp)

	return targetSumWaysTab(num, target)
}

// Recursion
// T.C = exponentially
func targetSumWaysRes(numbers []int, target int) int {

	// Base case
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}

	ans := 0
	for _, num := range numbers {
		ans = ans + targetSumWaysRes(numbers, target-num)
	}

	return ans
}

// Memoization
// T.C = O(n), S.C = O(n) + O(n)
func targetSumWaysDp(numbers []int, target int, dp []int) int {

	// Base case
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}

	if dp[target] != -1 {
		return dp[target]
	}

	ans := 0
	for _, num := range numbers {
		ans = ans + targetSumWaysDp(numbers, target-num, dp)
	}

	dp[target] = ans

	return dp[target]
}

// Tabulation
// T.C = O(n), S.C = O(n)
func targetSumWaysTab(numbers []int, target int) int {

	dp := slices.Repeat([]int{0}, target+1)

	// based on base case
	dp[0] = 1

	// traversing from target 1 to target
	for i := 1; i <= target; i++ {

		// traversing on num array
		for _, num := range numbers {
			if i-num >= 0 {
				dp[i] = dp[i] + dp[i-num]
			}
		}
	}

	return dp[target]
}
