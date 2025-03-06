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
