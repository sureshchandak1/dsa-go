package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMinOperationsToEqual1(t *testing.T) {
	fmt.Println(minOperationsToEqual1([]int{0, 1, 1, 1, 0, 0}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 1, 1}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 0, 1, 0}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 0, 0, 1, 0, 1}))
	fmt.Println(minOperationsToEqual1([]int{1, 1, 0, 1}))
}

func TestLongestNiceSubarray(t *testing.T) {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10}))
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))
}
