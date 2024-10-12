package stack

import (
	"fmt"
	"testing"
)

func TestNextSmallerElement(t *testing.T) {
	fmt.Println(nextSmallerElement([]int{2, 1, 4, 3}))
	fmt.Println(nextSmallerElement([]int{2, 1, 4, 2, 3}))
	fmt.Println(nextSmallerElement([]int{1, 3, 2}))
	fmt.Println(nextSmallerElement([]int{1, 2, 3, 4}))
}

func TestPrevSmallerElements(t *testing.T) {
	fmt.Println(prevSmallerElements([]int{2, 1, 4, 3}))
	fmt.Println(prevSmallerElements([]int{2, 1, 4, 2, 3}))
	fmt.Println(prevSmallerElements([]int{1, 3, 2}))
	fmt.Println(prevSmallerElements([]int{1, 2, 3, 4}))
}

func TestLargestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{1, 0, 1, 2, 2, 2, 2, 1, 0, 2}))
	fmt.Println(largestRectangleArea([]int{1, 2, 1, 0, 1, 1, 0, 0, 2, 2}))
	fmt.Println(largestRectangleArea([]int{8, 6, 3, 5, 0, 0, 4, 10, 2, 5}))
	fmt.Println(largestRectangleArea([]int{6, 1, 8, 10, 5, 7, 0, 4, 5, 8}))
	fmt.Println(largestRectangleArea([]int{2, 2, 2, 2, 2, 2, 2}))
	fmt.Println(largestRectangleArea([]int{2}))
	fmt.Println(largestRectangleArea([]int{2, 2}))
}
