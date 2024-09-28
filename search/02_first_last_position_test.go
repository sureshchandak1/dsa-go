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

func TestNumberOfOccurrence(t *testing.T) {

	arr := []int{1, 8, 12, 15, 17, 17, 18, 18, 18, 18, 19}

	fmt.Println("Number of Occurrence: ", numberOfOccurrence(arr, 17))
	fmt.Println("Number of Occurrence: ", numberOfOccurrence(arr, 18))
	fmt.Println("Number of Occurrence: ", numberOfOccurrence(arr, 19))
}
