package search

import (
	"fmt"
	"testing"
)

func TestPivot(t *testing.T) {

	arr := []int{8, 10, 17, 1, 3}

	index, element := pivot(arr)
	fmt.Printf("Index: %d, Element: %d", index, element)

}

func TestSearchRotatedArray(t *testing.T) {

	arr := []int{8, 10, 17, 1, 3}

	fmt.Println(searchRotatedArray(arr, 1))
	fmt.Println(searchRotatedArray(arr, 30))

}
