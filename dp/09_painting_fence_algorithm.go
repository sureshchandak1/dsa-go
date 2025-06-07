package dp

import "slices"

func numberOfWaysPaint(n int, k int) int {
	// return numberOfWaysPaintRec(n, k)

	// dp := slices.Repeat([]int{-1}, n+1)
	// return numberOfWaysPaintMem(n, k, dp)

	// return numberOfWaysPaintTab(n, k)

	return numberOfWaysPaintOptimize(n, k)
}

// Recursion
// T.C = exponentially
func numberOfWaysPaintRec(n int, k int) int {

	// Base case
	if n == 1 {
		return k
	}
	if n == 2 {
		return add(k, multiply(k, k-1))
	}

	first := multiply(numberOfWaysPaintRec(n-2, k), k-1)
	second := multiply(numberOfWaysPaintRec(n-1, k), k-1)

	ans := add(first, second)

	return ans
}

// Memoization
// T.C = O(n), S.C = O(n) + O(n)
func numberOfWaysPaintMem(n int, k int, dp []int) int {

	// Base case
	if n == 1 {
		return k
	}
	if n == 2 {
		return add(k, multiply(k, k-1))
	}

	if dp[n] != -1 {
		return dp[n]
	}

	first := multiply(numberOfWaysPaintMem(n-2, k, dp), k-1)
	second := multiply(numberOfWaysPaintMem(n-1, k, dp), k-1)

	dp[n] = add(first, second)

	return dp[n]
}

// Tabulation
// T.C = O(n), S.C = O(n)
func numberOfWaysPaintTab(n int, k int) int {

	if n == 1 {
		return k
	}

	dp := slices.Repeat([]int{0}, n+1)

	// according base case
	dp[1] = k
	dp[2] = add(k, multiply(k, k-1))

	for i := 3; i <= n; i++ {
		first := multiply(dp[i-2], k-1)
		second := multiply(dp[i-1], k-1)

		dp[i] = add(first, second)
	}

	return dp[n]

}

// Space Optimization
// T.C = O(n), S.C = O(1)
func numberOfWaysPaintOptimize(n int, k int) int {

	prev2 := k
	prev1 := add(k, multiply(k, k-1))

	for i := 3; i <= n; i++ {
		first := multiply(prev2, k-1)
		second := multiply(prev1, k-1)

		ans := add(first, second)

		prev2 = prev1
		prev1 = ans
	}

	return prev1
}

func add(a int, b int) int {
	return (a%MOD + b%MOD) % MOD
}

func multiply(a int, b int) int {
	return (a % MOD * b % MOD) % MOD
}
