package dp

import (
	"math"
	"slices"
)

func perfectSquares(n int) int {
	// return perfectSquaresRes(n)

	// dp := slices.Repeat([]int{-1}, n+1)
	// return perfectSquaresDp(n, dp)

	return perfectSquaresTab(n)
}

// Recursion
// T.C = exponentially
func perfectSquaresRes(n int) int {

	// Base case
	if n == 0 {
		return 0
	}

	ans := n
	i := 1

	for i*i <= n {
		ans = min(ans, 1+perfectSquaresRes(n-(i*i)))

		i++
	}

	return ans
}

// Memorization
// T.C = O(n), S.C = O(n) + O(n)
func perfectSquaresDp(n int, dp []int) int {

	// Base case
	if n == 0 {
		return 0
	}

	if dp[n] != -1 {
		return dp[n]
	}

	ans := n
	i := 1

	for i*i <= n {
		ans = min(ans, 1+perfectSquaresDp(n-(i*i), dp))

		i++
	}

	dp[n] = ans
	return dp[n]
}

// Tabulation
// T.C = O(n), S.C = O(n)
func perfectSquaresTab(n int) int {

	dp := slices.Repeat([]int{math.MaxInt}, n+1)

	// according based case
	dp[0] = 0

	for index := 1; index <= n; index++ {

		i := 1

		for i*i <= n {
			if index-(i*i) >= 0 {
				dp[index] = min(dp[index], 1+dp[index-(i*i)])
			}

			i++
		}
	}

	return dp[n]
}
