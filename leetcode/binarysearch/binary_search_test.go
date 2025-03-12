package binarysearch

import (
	"fmt"
	"testing"
)

func TestMaximumCount(t *testing.T) {
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
	fmt.Println(maximumCount([]int{-3, -2, -1, 0, 0, 1, 2}))
	fmt.Println(maximumCount([]int{5, 20, 66, 1314}))

	fmt.Println(maximumCountBruteForce([]int{-2, -1, -1, 1, 2, 3}))
	fmt.Println(maximumCountBruteForce([]int{-3, -2, -1, 0, 0, 1, 2}))
	fmt.Println(maximumCountBruteForce([]int{5, 20, 66, 1314}))
}
