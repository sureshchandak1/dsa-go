package dp

import "slices"

func equalPartition(arr []int) int {

	sum := 0
	for i := range arr {
		sum = sum + arr[i]
	}

	if sum%2 != 0 { // odd number
		return 0
	}

	// target := sum / 2

	// return equalPartitionRec(arr, 0, target)

	// dp := make([][]int, len(arr)+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, target+1)
	// }
	// return equalPartitionMem(arr, 0, target, dp)

	// return equalPartitionTab(arr, sum)

	return equalPartitionSC(arr, sum)
}

// Recursion
func equalPartitionRec(arr []int, index, target int) int {

	// Base case
	if index >= len(arr) {
		return 0
	}
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}

	include := equalPartitionRec(arr, index+1, target-arr[index])

	exclude := equalPartitionRec(arr, index+1, target)

	return include | exclude
}

// Recursion + Memorization
// T.C = O(n * total/2), S.C = O(n * total/2)
func equalPartitionMem(arr []int, index, target int, dp [][]int) int {

	// Base case
	if index >= len(arr) {
		return 0
	}
	if target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}

	if dp[index][target] != -1 {
		return dp[index][target]
	}

	include := equalPartitionMem(arr, index+1, target-arr[index], dp)

	exclude := equalPartitionMem(arr, index+1, target, dp)

	dp[index][target] = include | exclude

	return dp[index][target]
}

// Tabulation
// T.C = O(n * total/2), S.C = O(n * total/2)
func equalPartitionTab(arr []int, total int) int {

	n := len(arr)

	dp := make([][]int, len(arr)+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, total/2+1)
	}

	// based on base case
	for i := range dp {
		dp[i][0] = 1
	}

	for index := n - 1; index >= 0; index-- {
		for target := 0; target <= total/2; target++ {

			include := 0
			if target-arr[index] >= 0 {
				include = dp[index+1][target-arr[index]]
			}

			exclude := dp[index+1][target]

			dp[index][target] = include | exclude

		}
	}

	return dp[0][total/2]
}

// Space Optimization
// T.C = O(n * total/2), S.C = O(total/2)
func equalPartitionSC(arr []int, total int) int {

	n := len(arr)

	curr := make([]int, total/2+1)
	next := make([]int, total/2+1)

	// based on base case
	curr[0] = 1
	next[0] = 1

	for index := n - 1; index >= 0; index-- {
		for target := 0; target <= total/2; target++ {

			include := 0
			if target-arr[index] >= 0 {
				include = next[target-arr[index]]
			}

			exclude := next[target]

			curr[target] = include | exclude

		}

		next = slices.Clone(curr)
	}

	return next[total/2]
}
