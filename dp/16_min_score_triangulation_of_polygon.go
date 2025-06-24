package dp

import (
	"math"
	"slices"
)

func minScoreTriangulation(values []int) int {
	// n := len(values)
	// return minScoreTriRes(values, 0, n-1)

	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, n+1)
	// }
	// return minScoreTriDp(values, 0, n-1, dp)

	return minScoreTriTab(values)
}

// only Recursion
func minScoreTriRes(v []int, i, j int) int {

	// Base case
	if i+1 == j {
		return 0
	}

	ans := math.MaxInt
	temp := 0
	for k := i + 1; k < j; k++ {
		temp = v[i]*v[j]*v[k] + minScoreTriRes(v, i, k) + minScoreTriRes(v, k, j)
		ans = min(ans, temp)
	}

	return ans

}

// Recursion + Memorization
func minScoreTriDp(v []int, i, j int, dp [][]int) int {

	// Base case
	if i+1 == j {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	ans := math.MaxInt
	temp := 0
	for k := i + 1; k < j; k++ {
		temp = v[i]*v[j]*v[k] + minScoreTriDp(v, i, k, dp) + minScoreTriDp(v, k, j, dp)
		ans = min(ans, temp)
	}

	dp[i][j] = ans
	return dp[i][j]

}

// Tabulation
func minScoreTriTab(v []int) int {
	n := len(v)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {

			ans := math.MaxInt
			temp := 0
			for k := i + 1; k < j; k++ {
				temp = v[i]*v[j]*v[k] + dp[i][k] + dp[k][j]
				ans = min(ans, temp)
			}

			dp[i][j] = ans

		}
	}

	return dp[0][n-1]

}
