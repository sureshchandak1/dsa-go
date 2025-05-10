package dp

import "slices"

/**
 *   2 approaches
 *   1. Top-Down : Recursion + Memoization (Store the value of sub problems in map/table)
 *   2. Bottom-Up : Tabulation
 *
 *   Space Optimisation
 */

func fib(n int) int {

	// Base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	val1 := fib(n - 1)
	val2 := fib(n - 2)

	sum := val1 + val2

	return sum
}

func fibonacci(n int) int {

	dp := slices.Repeat([]int{-1}, n+1)

	fibonacciSolve(n, dp)

	fibonacciSolve2(n, dp)

	return fibonacciSolve3(n)
}

/**
 *   Top-Down
 *   T.C = O(n)
 *   S.C = O(n) + O(n)
 */
func fibonacciSolve(n int, dp []int) int {

	// Base case
	if n <= 1 {
		return n
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = fibonacciSolve(n-1, dp) + fibonacciSolve(n-2, dp)

	return dp[n]
}

/**
 *    Bottom-Up
 *    T.C = O(n)
 *    S.C = O(n)
 */
func fibonacciSolve2(n int, dp []int) int {

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

/**
 *    Space Optimisation
 *    T.C = O(n)
 *    S.C = O(1)
 */
func fibonacciSolve3(n int) int {

	prev1 := 0
	prev2 := 1
	curr := 0

	for i := 2; i <= n; i++ {
		curr = prev1 + prev2

		prev1 = prev2
		prev2 = curr
	}

	return prev2
}
