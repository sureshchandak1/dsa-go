package leetcode

import (
	"fmt"
	"testing"
)

func TestColoredCells(t *testing.T) {
	fmt.Println(coloredCells(3))
	fmt.Println(solveColoredCells(3))

	fmt.Println(coloredCells(5))
	fmt.Println(solveColoredCells(5))
}

func TestFindMissingAndRepeatedValues(t *testing.T) {
	grid := [][]int{
		{9, 1, 7},
		{8, 9, 2},
		{3, 4, 6},
	}
	fmt.Println(findMissingAndRepeatedValues(grid))
}
func TestClosestPrimes(t *testing.T) {
	fmt.Println(closestPrimes(10, 19))
	fmt.Println(closestPrimes(4, 6))
}

func TestMinimumRecolors(t *testing.T) {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}

func TestNumberOfAlternatingGroups(t *testing.T) {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 1, 0}, 3))
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1, 0, 1}, 6))
	fmt.Println(numberOfAlternatingGroups([]int{1, 1, 0, 1}, 4))
}

func TestCountOfSubstrings(t *testing.T) {
	fmt.Println(countOfSubstrings("aeioqq", 1))
	fmt.Println(countOfSubstrings("aeiou", 0))
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1))
}
