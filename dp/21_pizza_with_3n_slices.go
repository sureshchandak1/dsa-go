package dp

import "slices"

func maxSizeSlices(slice []int) int {

	// k := len(slice)
	// case1 := maxSizeSlicesRec(0, k-2, slice, k/3)
	// case2 := maxSizeSlicesRec(1, k-1, slice, k/3)
	// return max(case1, case2)

	// dp1 := make([][]int, k)
	// for i := range dp1 {
	// 	dp1[i] = slices.Repeat([]int{-1}, k)
	// }
	// dp2 := make([][]int, k)
	// for i := range dp2 {
	// 	dp2[i] = slices.Repeat([]int{-1}, k)
	// }
	// case1 := maxSizeSlicesMem(0, k-2, slice, k/3, dp1)
	// case2 := maxSizeSlicesMem(1, k-1, slice, k/3, dp2)

	// return max(case1, case2)

	// return maxSizeSlicesTab(slice)

	return maxSizeSlicesSC(slice)
}

// Recursion
func maxSizeSlicesRec(index, endIndex int, slices []int, totalEat int) int {

	// Base case
	if totalEat == 0 || index > endIndex {
		return 0
	}

	// take slice
	include := slices[index] + maxSizeSlicesRec(index+2, endIndex, slices, totalEat-1)

	// not take slice
	exclude := maxSizeSlicesRec(index+1, endIndex, slices, totalEat)

	return max(include, exclude)

}

// Recursion + Memorization
func maxSizeSlicesMem(index, endIndex int, slices []int, totalEat int, dp [][]int) int {

	// Base case
	if totalEat == 0 || index > endIndex {
		return 0
	}

	if dp[index][totalEat] != -1 {
		return dp[index][totalEat]
	}

	// take slice
	include := slices[index] + maxSizeSlicesMem(index+2, endIndex, slices, totalEat-1, dp)

	// not take slice
	exclude := maxSizeSlicesMem(index+1, endIndex, slices, totalEat, dp)

	dp[index][totalEat] = max(include, exclude)
	return dp[index][totalEat]

}

// Tabulation
func maxSizeSlicesTab(slice []int) int {

	k := len(slice)
	dp1 := make([][]int, k+2)
	for i := range dp1 {
		dp1[i] = slices.Repeat([]int{0}, k+2)
	}
	dp2 := make([][]int, k+2)
	for i := range dp2 {
		dp2[i] = slices.Repeat([]int{0}, k+2)
	}

	for index := k - 2; index >= 0; index-- {

		for totalEat := 1; totalEat <= k/3; totalEat++ {

			// take slice
			include := slice[index] + dp1[index+2][totalEat-1]

			// not take slice
			exclude := dp1[index+1][totalEat]

			dp1[index][totalEat] = max(include, exclude)
		}
	}
	case1 := dp1[0][k/3]

	for index := k - 1; index >= 1; index-- {

		for totalEat := 1; totalEat <= k/3; totalEat++ {

			// take slice
			include := slice[index] + dp2[index+2][totalEat-1]

			// not take slice
			exclude := dp2[index+1][totalEat]

			dp2[index][totalEat] = max(include, exclude)
		}
	}
	case2 := dp2[1][k/3]

	return max(case1, case2)

}

// Space Optimization
func maxSizeSlicesSC(slice []int) int {

	k := len(slice)

	prev1 := slices.Repeat([]int{0}, k+2)
	curr1 := slices.Repeat([]int{0}, k+2)
	next1 := slices.Repeat([]int{0}, k+2)

	prev2 := slices.Repeat([]int{0}, k+2)
	curr2 := slices.Repeat([]int{0}, k+2)
	next2 := slices.Repeat([]int{0}, k+2)

	for index := k - 2; index >= 0; index-- {

		for totalEat := 1; totalEat <= k/3; totalEat++ {

			// take slice
			include := slice[index] + next1[totalEat-1]

			// not take slice
			exclude := curr1[totalEat]

			prev1[totalEat] = max(include, exclude)
		}

		next1 = slices.Clone(curr1)
		curr1 = slices.Clone(prev1)
	}
	case1 := curr1[k/3]

	for index := k - 1; index >= 1; index-- {

		for totalEat := 1; totalEat <= k/3; totalEat++ {

			// take slice
			include := slice[index] + next2[totalEat-1]

			// not take slice
			exclude := curr2[totalEat]

			prev2[totalEat] = max(include, exclude)
		}

		next2 = slices.Clone(curr2)
		curr2 = slices.Clone(prev2)
	}
	case2 := curr2[k/3]

	return max(case1, case2)
}
