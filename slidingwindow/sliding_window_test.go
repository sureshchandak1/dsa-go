package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMaxSumSubarray(t *testing.T) {
	fmt.Println(maxSumSubarray([]int{100, 200, 300, 400}, 2))
	fmt.Println(maxSumSubarray([]int{1, 5, 4, 2, 9, 9, 9}, 3))
}

func TestMinimumRecolors(t *testing.T) {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}

func TestMaximumSubarraySum(t *testing.T) {
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
}

func TestFirstNegativeNumberK(t *testing.T) {
	fmt.Println(firstNegativeNumberKBrutForce([]int{-8, 2, 3, -6, 10}, 2))
	fmt.Println(firstNegativeNumberKBrutForce([]int{12, -1, -7, 8, -15, 30, 16, 28}, 3))
	fmt.Println(firstNegativeNumberKBrutForce([]int{12, 1, 3, 5}, 3))

	fmt.Println(firstNegativeNumberK([]int{-8, 2, 3, -6, 10}, 2))
	fmt.Println(firstNegativeNumberK([]int{12, -1, -7, 8, -15, 30, 16, 28}, 3))
	fmt.Println(firstNegativeNumberK([]int{12, 1, 3, 5}, 3))
}

func TestChocolateDistribution(t *testing.T) {
	fmt.Println(chocolateDistribution([]int{3, 4, 1, 9, 56, 7, 9, 12}, 5))
	fmt.Println(chocolateDistribution([]int{1, 5, 4, 2, 9, 1, 2}, 3))
}

func TestSmallestSubarrayGreaterX(t *testing.T) {
	fmt.Println(smallestSubarrayGreaterX([]int{1, 4, 45, 6, 0, 19}, 51))
	fmt.Println(smallestSubarrayGreaterX([]int{3, 4, 1, 9, 56, 7, 9, 12}, 15))
	fmt.Println(smallestSubarrayGreaterX([]int{1, 5, 4, 2, 9, 1, 2}, 13))
}

func TestLongestKUniqueSubStr(t *testing.T) {
	fmt.Println(longestKUniqueSubStr("aabacbebebe", 3))
	fmt.Println(longestKUniqueSubStr("sgfasafsfdsfsfaffa", 2))
	fmt.Println(longestKUniqueSubStr("fhdfgerthgasdgas", 3))
	fmt.Println(longestKUniqueSubStr("", 3))
}
