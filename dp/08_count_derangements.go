package dp

import "slices"

var MOD int = 1000000007

func countDerangements(n int) int {

	// return countDerangementsRec(n)

	// dp := slices.Repeat([]int{-1}, n+1)
	// return countDerangementsMem(n, dp)

	// return countDerangementsTab(n)

	return countDerangementsOP(n)

}

// Recursion
// T.C = exponentially
func countDerangementsRec(n int) int {

	// Base case
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	ans := (n - 1) % MOD * (countDerangementsRec(n-1)%MOD + countDerangementsRec(n-2)%MOD) % MOD

	return ans

}

// Memoization
// T.C = O(n), S.C = O(n) + O(n)
func countDerangementsMem(n int, dp []int) int {

	// Base case
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = (n - 1) % MOD * (countDerangementsMem(n-1, dp)%MOD + countDerangementsMem(n-2, dp)%MOD) % MOD

	return dp[n]

}

// Tabulation
// T.C = O(n), S.C = O(n)
func countDerangementsTab(n int) int {

	dp := slices.Repeat([]int{0}, n+1)

	// according to base case
	dp[1] = 0
	if n >= 2 {
		dp[2] = 1
	}

	for i := 3; i <= n; i++ {
		first := dp[i-1] % MOD
		second := dp[i-2] % MOD
		sum := (first + second) % MOD

		ans := ((i - 1) * sum) % MOD

		dp[i] = ans
	}

	return dp[n]
}

// Space Optimization
// T.C = O(n), S.C = O(1)
func countDerangementsOP(n int) int {

	if n == 1 {
		return 0
	}

	prev1 := 1
	prev2 := 0

	for i := 3; i <= n; i++ {
		first := prev1 % MOD
		second := prev2 % MOD
		sum := (first + second) % MOD

		ans := ((i - 1) * sum) % MOD

		prev2 = prev1
		prev1 = ans
	}

	return prev1
}
