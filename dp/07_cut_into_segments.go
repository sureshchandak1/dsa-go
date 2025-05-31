package dp

import (
	"math"
	"slices"
)

func cutSegments(n, x, y, z int) int {
	// ans := cutSegmentsRec(n, x, y, z)

	// dp := slices.Repeat([]int{-1}, n+1)
	// ans := cutSegmentsMem(n, x, y, z, dp)

	// if ans < 0 {
	// 	return 0
	// }
	// return ans

	return cutSegmentsTab(n, x, y, z)
}

// T.C = infinite
func cutSegmentsRec(n, x, y, z int) int {

	// Base case
	if n == 0 {
		return 0
	}
	if n < 0 {
		return math.MinInt
	}

	xAns := cutSegmentsRec(n-x, x, y, z) + 1
	yAns := cutSegmentsRec(n-y, x, y, z) + 1
	zAns := cutSegmentsRec(n-z, x, y, z) + 1

	ans := max(xAns, max(yAns, zAns))

	return ans
}

// Memoization
// T.C = O(n), S.C = O(n) + O(n)
func cutSegmentsMem(n, x, y, z int, dp []int) int {

	// Base case
	if n == 0 {
		return 0
	}
	if n < 0 {
		return math.MinInt
	}

	if dp[n] != -1 {
		return dp[n]
	}

	xAns := cutSegmentsMem(n-x, x, y, z, dp) + 1
	yAns := cutSegmentsMem(n-y, x, y, z, dp) + 1
	zAns := cutSegmentsMem(n-z, x, y, z, dp) + 1

	ans := max(xAns, max(yAns, zAns))

	dp[n] = ans

	return dp[n]
}

// Tabulation
// T.C = O(n), S.C = O(n)
func cutSegmentsTab(n, x, y, z int) int {

	dp := slices.Repeat([]int{math.MinInt}, n+1)

	// based on base case
	dp[0] = 0

	for i := 1; i <= n; i++ {

		if i-x >= 0 {
			dp[i] = max(dp[i], dp[i-x]+1)
		}

		if i-y >= 0 {
			dp[i] = max(dp[i], dp[i-y]+1)
		}

		if i-z >= 0 {
			dp[i] = max(dp[i], dp[i-z]+1)
		}

	}

	if dp[n] < 0 {
		return 0
	}

	return dp[n]
}
