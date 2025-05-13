package dp

func climbStair(nStairs int, position int) int {

	// base case
	if position == nStairs {
		return 1
	}
	if position > nStairs {
		return 0
	}

	return climbStair(nStairs, position+1) + climbStair(nStairs, position+2)
}

func minCostClimbingStairs(cost []int) int {

	n := len(cost)

	// ans1 := min(solveMinCostClimbingStairs(cost, n-1), solveMinCostClimbingStairs(cost, n-2))

	// dp := slices.Repeat([]int{-1}, n+1)
	// ans2 := min(solveMinCostClimbingStairs1(cost, n-1, dp), solveMinCostClimbingStairs1(cost, n-2, dp))

	// return solveMinCostClimbingStairs2(cost, n)

	return solveMinCostClimbingStairs3(cost, n)
}

func solveMinCostClimbingStairs(cost []int, n int) int {

	// Base case
	if n == 0 {
		return cost[0]
	}
	if n == 1 {
		return cost[1]
	}

	ans := cost[n] + min(solveMinCostClimbingStairs(cost, n-1), solveMinCostClimbingStairs(cost, n-2))

	return ans
}

// with dp: Recursion + Memoization
// T.C = O(n), S.C = O(n) + O(n)
func solveMinCostClimbingStairs1(cost []int, n int, dp []int) int {

	// Base case
	if n == 0 {
		return cost[0]
	}
	if n == 1 {
		return cost[1]
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = cost[n] + min(solveMinCostClimbingStairs1(cost, n-1, dp), solveMinCostClimbingStairs1(cost, n-2, dp))

	return dp[n]
}

// Tabulation(Bottom-Up)
// T.C = O(n), S.C = O(n)
func solveMinCostClimbingStairs2(cost []int, n int) int {

	dp := make([]int, n+1)

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}

// optimized solution: Space Optimisation
// T.C = O(n), S.C = O(1)
func solveMinCostClimbingStairs3(cost []int, n int) int {

	prev1 := cost[0]
	prev2 := cost[1]
	curr := 0

	for i := 2; i < n; i++ {
		curr = cost[i] + min(prev1, prev2)

		prev1 = prev2
		prev2 = curr
	}

	return min(prev1, prev2)
}
