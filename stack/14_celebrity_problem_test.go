package stack

import (
	"fmt"
	"testing"
)

func TestCelebrityProblem(t *testing.T) {

	matrix := [][]int{
		{0, 1, 0},
		{0, 0, 0},
		{0, 1, 0},
	}

	fmt.Println(celebrityProblem(matrix))

	matrix = [][]int{
		{0, 1},
		{1, 0},
	}

	fmt.Println(celebrityProblem(matrix))
}
