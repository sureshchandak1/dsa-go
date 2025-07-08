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

func TestCutSegments(t *testing.T) {
	fmt.Println(cutSegments(7, 5, 2, 2))
	fmt.Println(cutSegments(8, 3, 3, 3))
	fmt.Println(cutSegments(7, 3, 2, 2))
	fmt.Println(cutSegments(8, 1, 4, 4))
}

func TestCountDerangements(t *testing.T) {
	fmt.Println(countDerangements(2))
	fmt.Println(countDerangements(3))
	fmt.Println(countDerangements(1))
	fmt.Println(countDerangements(4))
}

func TestNumberOfWaysPaint(t *testing.T) {
	fmt.Println(numberOfWaysPaint(1, 1))
	fmt.Println(numberOfWaysPaint(3, 2))
	fmt.Println(numberOfWaysPaint(2, 4))
	fmt.Println(numberOfWaysPaint(4, 2))
}

func TestKnapsack(t *testing.T) {
	var weight = []int{25, 4, 25, 49, 9, 11, 31, 5, 37, 7, 11, 47, 37, 1, 33, 25, 29, 11, 1, 41, 19, 14, 43, 21, 1, 21, 23, 37, 12, 11, 45, 13, 36, 11, 17}
	var value = []int{86, 55, 17, 31, 88, 82, 27, 57, 18, 1, 61, 56, 36, 82, 51, 85, 55, 21, 11, 76, 91, 36, 85, 32, 99, 43, 41, 61, 41, 28, 83, 27, 35, 61, 16}
	fmt.Println(knapsack(weight, value, len(weight), 869))

	weight = []int{30, 3, 43, 25, 20, 21, 2, 34, 39, 48, 14, 35, 15, 46, 49, 15, 11, 26, 38, 33, 1, 21, 17, 1, 44, 44, 33, 17, 33}
	value = []int{1, 13, 75, 46, 87, 98, 31, 91, 23, 43, 85, 18, 35, 1, 83, 26, 77, 91, 85, 76, 74, 65, 35, 36, 98, 42, 37, 83, 54}
	fmt.Println(knapsack(weight, value, len(weight), 424))
}

func TestTargetSumWays(t *testing.T) {
	fmt.Println(targetSumWays([]int{1, 2, 3}, 4))
	fmt.Println(targetSumWays([]int{9}, 3))
	fmt.Println(targetSumWays([]int{1, 2, 5}, 5))
	fmt.Println(targetSumWays([]int{1, 2, 5}, 6))
}

func TestPerfectSquares(t *testing.T) {
	fmt.Println(perfectSquares(12))
	fmt.Println(perfectSquares(13))
	fmt.Println(perfectSquares(10000))
}

func TestMinCostTickets(t *testing.T) {
	fmt.Println(minCostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(minCostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
	fmt.Println()
	fmt.Println(minCostTicketsOptimized([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(minCostTicketsOptimized([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}

func TestMaxSquare(t *testing.T) {
	fmt.Println(maxSquare([][]int{{1, 1}, {1, 1}}))
	fmt.Println(maxSquare([][]int{{0, 0}, {0, 0}}))
	fmt.Println(maxSquare([][]int{{0, 1, 1, 0, 1}, {1, 1, 0, 1, 0}, {0, 1, 1, 1, 0}, {1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}}))
}

func TestMinScoreTriangulation(t *testing.T) {
	fmt.Println(minScoreTriangulation([]int{1, 2, 3}))
	fmt.Println(minScoreTriangulation([]int{3, 7, 4, 5}))
	fmt.Println(minScoreTriangulation([]int{1, 3, 1, 4, 1, 5}))
}

func TestMinSideJumps(t *testing.T) {
	fmt.Println(minSideJumps([]int{0, 1, 2, 3, 0}))
	fmt.Println(minSideJumps([]int{0, 1, 1, 3, 3, 0}))
	fmt.Println(minSideJumps([]int{0, 2, 1, 0, 3, 0}))
}

func TestMaxSatisfaction(t *testing.T) {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -7}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))

}
