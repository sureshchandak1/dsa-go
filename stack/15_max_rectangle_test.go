package stack

import (
	"fmt"
	"testing"
)

func TestMaximalRectangle(t *testing.T) {
	fmt.Println(maximalRectangle([][]int{
		{0, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 0, 0},
	}))

	fmt.Println(maximalRectangle([][]int{
		{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 0},
	}))

}
