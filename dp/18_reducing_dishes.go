package dp

import (
	"slices"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {

	sort.Ints(satisfaction)

	// return maxSatisfactionRes(satisfaction, 0, 0)

	// n := len(satisfaction)
	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, n+1)
	// }

	// return maxSatisfactionDp(satisfaction, 0, 0, dp)

	// return maxSatisfactionTab(satisfaction)

	return maxSatisfactionSC(satisfaction)
}

// Recursion
func maxSatisfactionRes(satisfaction []int, index, time int) int {

	// base case
	if index == len(satisfaction) {
		return 0
	}

	include := (time+1)*satisfaction[index] + maxSatisfactionRes(satisfaction, index+1, time+1)

	exclude := maxSatisfactionRes(satisfaction, index+1, time)

	return max(include, exclude)
}

// Recursion + Memorization
// T.C = O(n * n), S.C = O(n * n)
func maxSatisfactionDp(satisfaction []int, index, time int, dp [][]int) int {

	// base case
	if index == len(satisfaction) {
		return 0
	}

	if dp[index][time] != -1 {
		return dp[index][time]
	}

	include := (time+1)*satisfaction[index] + maxSatisfactionDp(satisfaction, index+1, time+1, dp)

	exclude := maxSatisfactionDp(satisfaction, index+1, time, dp)

	dp[index][time] = max(include, exclude)

	return dp[index][time]
}

// Tabulation
// T.C = O(n * n), S.C = O(n * n)
func maxSatisfactionTab(satisfaction []int) int {

	n := len(satisfaction)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, n+1)
	}

	for index := n - 1; index >= 0; index-- {

		for time := index; time >= 0; time-- {

			include := (time+1)*satisfaction[index] + dp[index+1][time+1]

			exclude := dp[index+1][time]

			dp[index][time] = max(include, exclude)

		}
	}

	return dp[0][0]

}

// Space Optimization
// T.C = O(n * n), S.C = O(n)
func maxSatisfactionSC(satisfaction []int) int {

	n := len(satisfaction)

	curr := make([]int, n+1)
	next := make([]int, n+1)

	for index := n - 1; index >= 0; index-- {

		for time := index; time >= 0; time-- {

			include := (time+1)*satisfaction[index] + next[time+1]

			exclude := next[time]

			curr[time] = max(include, exclude)

		}

		next = slices.Clone(curr)
	}

	return next[0]
}
