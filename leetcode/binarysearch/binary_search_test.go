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

func TestMinCapability(t *testing.T) {
	fmt.Println(minCapability([]int{2, 3, 5, 9}, 2))
	fmt.Println(minCapability([]int{2, 7, 9, 3, 1}, 2))
}
