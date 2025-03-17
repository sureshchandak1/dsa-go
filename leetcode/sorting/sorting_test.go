package sorting

import (
	"fmt"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	fmt.Println(divideArraySolution1([]int{3, 2, 3, 2, 2, 2}))
	fmt.Println(divideArraySolution1([]int{1, 2, 3, 4}))
	fmt.Println()
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
	fmt.Println(divideArray([]int{1, 2, 3, 4}))
}
