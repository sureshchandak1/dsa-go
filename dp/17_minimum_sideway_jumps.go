package dp

import (
	"math"
	"slices"
)

func minSideJumps(obstacles []int) int {
	// return minSideJumpsRes(obstacles, 2, 0)

	// n := len(obstacles)
	// dp := make([][]int, 4)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, n)
	// }
	// return minSideJumpsDp(obstacles, 2, 0, dp)

	// return minSideJumpsTab(obstacles)

	return minSideJumpsSC(obstacles)
}

// Recursion
func minSideJumpsRes(obstacles []int, currLane, currPos int) int {

	// base case
	if currPos == len(obstacles)-1 {
		return 0
	}

	if obstacles[currPos+1] != currLane {
		return minSideJumpsRes(obstacles, currLane, currPos+1)
	} else {
		// sideways jump
		ans := math.MaxInt
		temp := 0
		for i := 1; i <= 3; i++ {
			if currLane != i && obstacles[currPos] != i {
				temp = 1 + minSideJumpsRes(obstacles, i, currPos)
				ans = min(ans, temp)
			}
		}
		return ans
	}

}

// Recursion + Memorization
// T.C = O(3 * n), S.C = O(4 * 2)
func minSideJumpsDp(obstacles []int, currLane, currPos int, dp [][]int) int {

	// base case
	if currPos == len(obstacles)-1 {
		return 0
	}

	if dp[currLane][currPos] != -1 {
		return dp[currLane][currPos]
	}

	if obstacles[currPos+1] != currLane {
		return minSideJumpsDp(obstacles, currLane, currPos+1, dp)
	} else {
		// sideways jump
		ans := math.MaxInt
		temp := 0
		for i := 1; i <= 3; i++ {
			if currLane != i && obstacles[currPos] != i {
				temp = 1 + minSideJumpsDp(obstacles, i, currPos, dp)
				ans = min(ans, temp)
			}
		}

		dp[currLane][currPos] = ans
		return dp[currLane][currPos]
	}

}

// Tabulation
// T.C = O(2 * n), S.C = O(4 * n)
func minSideJumpsTab(obstacles []int) int {

	n := len(obstacles)
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = slices.Repeat([]int{math.MaxInt}, n+1)
	}

	// Based on base case
	dp[0][n-1] = 0
	dp[1][n-1] = 0
	dp[2][n-1] = 0
	dp[3][n-1] = 0

	for currPos := n - 2; currPos >= 0; currPos-- {

		for currLane := 1; currLane <= 3; currLane++ {

			if obstacles[currPos+1] != currLane {
				dp[currLane][currPos] = dp[currLane][currPos+1]
			} else {
				// sideways jump
				ans := math.MaxInt
				temp := 0
				for i := 1; i <= 3; i++ {
					if currLane != i && obstacles[currPos] != i {
						temp = 1 + dp[i][currPos+1]
						ans = min(ans, temp)
					}
				}

				dp[currLane][currPos] = ans

			}
		}
	}

	return min(dp[2][0], min(1+dp[1][0], 1+dp[3][0]))
}

// Space optimization
// T.C = O(2 * n), S.C = O(1)
func minSideJumpsSC(obstacles []int) int {

	n := len(obstacles)

	curr := slices.Repeat([]int{math.MaxInt}, 4)
	next := slices.Repeat([]int{math.MaxInt}, 4)

	// Based on base case
	next[0] = 0
	next[1] = 0
	next[2] = 0
	next[3] = 0

	for currPos := n - 2; currPos >= 0; currPos-- {

		for currLane := 1; currLane <= 3; currLane++ {

			if obstacles[currPos+1] != currLane {
				curr[currLane] = next[currLane]
			} else {
				// sideways jump
				ans := math.MaxInt
				temp := 0
				for i := 1; i <= 3; i++ {
					if currLane != i && obstacles[currPos] != i {
						temp = 1 + next[i]
						ans = min(ans, temp)
					}
				}

				curr[currLane] = ans

			}
		}

		next = slices.Clone(curr)
	}

	return min(next[2], min(1+next[1], 1+next[3]))

}
