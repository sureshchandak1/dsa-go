package dp

import "slices"

func knapsack(weight, value []int, n, maxWeight int) int {
	// return knapsackRec(weight, value, n-1, maxWeight)

	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, maxWeight+1)
	// }
	// return knapsackDp(weight, value, n-1, maxWeight, dp)

	// return knapsackTab(weight, value, n, maxWeight)

	return knapsackSpaceOptimize(weight, value, n, maxWeight)
}

func knapsackRec(weight, value []int, index, maxWeight int) int {

	// Base case
	if index == 0 {
		if weight[index] <= maxWeight {
			return value[index]
		}
		return 0
	}

	// include value
	var include int = 0
	if weight[index] <= maxWeight {
		include = value[index] + knapsackRec(weight, value, index-1, maxWeight-weight[index])
	}

	// not include value
	exclude := knapsackRec(weight, value, index-1, maxWeight)

	return max(include, exclude)
}

// Top-Down DP (Memoization)
func knapsackDp(weight, value []int, index, maxWeight int, dp [][]int) int {

	// Base case
	if index == 0 {
		if weight[index] <= maxWeight {
			return value[index]
		}
		return 0
	}

	if dp[index][maxWeight] != -1 {
		return dp[index][maxWeight]
	}

	// include value
	var include int = 0
	if weight[index] <= maxWeight {
		include = value[index] + knapsackDp(weight, value, index-1, maxWeight-weight[index], dp)
	}

	// not include value
	exclude := knapsackDp(weight, value, index-1, maxWeight, dp)

	dp[index][maxWeight] = max(include, exclude)

	return dp[index][maxWeight]
}

func knapsackTab(weight, value []int, n, maxWeight int) int {

	// step1: create dp array
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, maxWeight+1)
	}

	// step2: set base case value
	for w := weight[0]; w <= maxWeight; w++ {
		if weight[0] <= maxWeight {
			dp[0][w] = value[0]
		} else {
			dp[0][w] = 0
		}
	}

	// step3: remaining part
	for index := 1; index < n; index++ {
		for w := 0; w <= maxWeight; w++ {

			// include value
			var include int = 0
			if weight[index] <= w {
				include = value[index] + dp[index-1][w-weight[index]]
			}

			// not include value
			exclude := dp[index-1][w]

			dp[index][w] = max(include, exclude)

		}
	}

	return dp[n-1][maxWeight]
}

func knapsackSpaceOptimize(weight, value []int, n, maxWeight int) int {

	// step1:
	prev := make([]int, maxWeight+1)
	curr := make([]int, maxWeight+1)

	// step2: set base case value
	for w := weight[0]; w <= maxWeight; w++ {
		if weight[0] <= maxWeight {
			prev[w] = value[0]
		} else {
			prev[w] = 0
		}
	}

	// step3: remaining part
	for index := 1; index < n; index++ {
		for w := 0; w <= maxWeight; w++ {

			// include value
			var include int = 0
			if weight[index] <= w {
				include = value[index] + prev[w-weight[index]]
			}

			// not include value
			exclude := prev[w]

			curr[w] = max(include, exclude)

		}

		prev = slices.Clone(curr)
	}

	return prev[maxWeight]
}
