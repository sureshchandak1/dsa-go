package dp

import (
	"math"
	"slices"
)

func coinChange(coins []int, amount int) int {
	// ans := coinChangeSolve(coins, amount)

	// dp := slices.Repeat([]int{-1}, amount+1)
	// ans := coinChangeSolveDP(coins, amount, dp)

	// if ans == math.MaxInt32 {
	// 	return -1
	// }

	return coinChangeSolveTab(coins, amount)
}

func coinChangeSolveTab(coins []int, amount int) int {

	dp := slices.Repeat([]int{math.MaxInt32}, amount+1)

	dp[0] = 0

	for i := 1; i <= amount; i++ {

		for j := 0; j < len(coins); j++ {

			if i-coins[j] >= 0 && dp[i-coins[j]] != math.MaxInt32 {
				dp[i] = min(dp[i], 1+dp[i-coins[j]])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]

}

func coinChangeSolveDP(coins []int, amount int, dp []int) int {

	// Base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt32
	}

	if dp[amount] != -1 {
		return dp[amount]
	}

	mini := math.MaxInt32
	ans := 0
	for coin := range coins {
		ans = coinChangeSolveDP(coins, amount-coin, dp)

		if ans != math.MaxInt32 {
			mini = min(mini, 1+ans)
		}
	}

	dp[amount] = mini

	return mini
}

func coinChangeSolve(coins []int, amount int) int {

	// Base case
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return math.MaxInt64
	}

	mini := math.MaxInt64
	ans := 0
	for coin := range coins {
		ans = coinChangeSolve(coins, amount-coin)

		if ans != math.MaxInt64 {
			mini = min(mini, 1+ans)
		}
	}

	return mini
}
