package search

import (
	"fmt"
	"testing"
)

func TestFirstLastIndex(t *testing.T) {

	arr := []int{0, 0, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4, 5, 6}

	firstIndex, lastIndex, element := firstLastPosition(arr, 1)
	fmt.Printf("%d: First: %d, Last: %d\n", element, firstIndex, lastIndex)

	firstIndex, lastIndex, element = firstLastPosition(arr, 3)
	fmt.Printf("%d: First: %d, Last: %d\n", element, firstIndex, lastIndex)

	firstIndex, lastIndex, element = firstLastPosition(arr, 5)
	fmt.Printf("%d: First: %d, Last: %d\n", element, firstIndex, lastIndex)
}
