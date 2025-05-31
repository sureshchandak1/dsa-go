package dp

import (
	"fmt"
	"testing"
)

func TestFibonacciSeries(t *testing.T) {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fibonacci(5))
}

func TestMinCostClimbingStairs(t *testing.T) {
	fmt.Println(climbStair(4, 0))
	fmt.Println(climbStair(6, 0))
	fmt.Println()
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func TestMinimumNumberCoins(t *testing.T) {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}

func TestMaxSumNonAdjacentElements(t *testing.T) {
	fmt.Println(maxSumNonAdjacentElements([]int{1, 2, 5, 4, 5, 6, 9, 10}))
}

func TestRobMoney1(t *testing.T) {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{1, 2}))

	fmt.Println()

	fmt.Println(rob2([]int{1, 2, 3, 1}))
	fmt.Println(rob2([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob2([]int{1, 2}))
}
