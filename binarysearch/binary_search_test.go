package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearchInc(t *testing.T) {
	fmt.Println(binarySearchOrderAgnostic([]int{1, 5, 11, 15, 18, 20, 21, 21, 22}, 22))
	fmt.Println(binarySearchOrderAgnostic([]int{22, 21, 21, 20, 18, 15, 11, 5, 1}, 22))
	fmt.Println(binarySearchOrderAgnostic([]int{2, 2, 2, 2, 2, 2, 2, 2, 2}, 2))
	fmt.Println(binarySearchOrderAgnostic([]int{1, 5, 11, 15, 18, 20, 21, 21, 22}, 50))
}
