package dp

import (
	"math"
	"slices"
)

func minCostTickets(days, cost []int) int {
	n := len(days)
	// return minCostTicketsRes(n, days, cost, 0)

	// dp := slices.Repeat([]int{-1}, n+1)
	// return minCostTicketsMem(n, days, cost, 0, dp)

	return minCostTicketsTab(n, days, cost)
}

func minCostTicketsRes(n int, days, cost []int, index int) int {

	// Base case
	if index >= n {
		return 0
	}

	// 1 day pass
	option1 := cost[0] + minCostTicketsRes(n, days, cost, index+1)

	// 7 day pass
	i := index
	for i < n && (days[i] < days[index]+7) {
		i++
	}

	option2 := cost[1] + minCostTicketsRes(n, days, cost, i)

	// 30 days pass
	i = index
	for i < n && (days[i] < days[index]+30) {
		i++
	}

	option3 := cost[2] + minCostTicketsRes(n, days, cost, i)

	return min(option1, min(option2, option3))
}

// Memorization
// T.C = O(n)
func minCostTicketsMem(n int, days, cost []int, index int, dp []int) int {

	// Base case
	if index >= n {
		return 0
	}

	if dp[index] != -1 {
		return dp[index]
	}

	// 1 day pass
	option1 := cost[0] + minCostTicketsMem(n, days, cost, index+1, dp)

	// 7 day pass
	i := index
	for i < n && (days[i] < days[index]+7) {
		i++
	}

	option2 := cost[1] + minCostTicketsMem(n, days, cost, i, dp)

	// 30 days pass
	i = index
	for i < n && (days[i] < days[index]+30) {
		i++
	}

	option3 := cost[2] + minCostTicketsMem(n, days, cost, i, dp)

	dp[index] = min(option1, min(option2, option3))

	return dp[index]
}

// T.C = O(n)
func minCostTicketsTab(n int, days, cost []int) int {

	dp := slices.Repeat([]int{math.MaxInt16}, n+1)

	// according base case
	dp[n] = 0

	for index := n - 1; index >= 0; index-- {

		// 1 day pass
		option1 := cost[0] + dp[index+1]

		// 7 day pass
		i := index
		for i < n && (days[i] < days[index]+7) {
			i++
		}

		option2 := cost[1] + dp[i]

		// 30 days pass
		i = index
		for i < n && (days[i] < days[index]+30) {
			i++
		}

		option3 := cost[2] + dp[i]

		dp[index] = min(option1, min(option2, option3))
	}

	return dp[0]
}
