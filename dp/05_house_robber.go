package dp

func rob(nums []int) int {
	// n := len(nums)
	// return robMoneyRec(nums, n-1)

	// dp := slices.Repeat([]int{-1}, n+1)
	// return robMoneyMem(nums, n-1, dp)

	return robMoneyTab(nums)
}

// Recursion
func robMoneyRec(nums []int, n int) int {

	// Base case
	if n < 0 {
		return 0
	}
	if n == 0 {
		return nums[0]
	}

	// include current element
	include := robMoneyRec(nums, n-2) + nums[n]

	// exclude current element
	exclude := robMoneyRec(nums, n-1)

	return max(include, exclude)
}

// Recursion + Memoization
// T.C = O(n), S.C = O(n) + O(n)
func robMoneyMem(nums []int, n int, dp []int) int {

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

	// include current element
	include := robMoneyMem(nums, n-2, dp) + nums[n]

	// exclude current element
	exclude := robMoneyMem(nums, n-1, dp)

	dp[n] = max(include, exclude)

	return dp[n]
}

func robMoneyTab(nums []int) int {
	n := len(nums)

	prev1 := nums[0]
	prev2 := 0

	for i := 1; i < n; i++ {
		incl := prev2 + nums[i]
		excl := prev1

		ans := max(incl, excl)

		prev2 = prev1
		prev1 = ans
	}

	return prev1
}
