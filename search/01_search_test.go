package search

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {

	arr := []int{5, 7, -2, 10, 22, -2, 0, 5, 22, 1}

	result, index, value := linearSearch(arr, 1)
	fmt.Printf("%t Index: %d, Value: %d\n", result, index, value)

	result, index, value = linearSearch(arr, 30)
	fmt.Printf("%t Index: %d, Value: %d\n", result, index, value)
}

func TestBinarySearch(t *testing.T) {

	arr := []int{2, 4, 6, 8, 12, 18}

	result, index, value := binarySearch(arr, 12)
	fmt.Printf("%t Index: %d, Value: %d\n", result, index, value)

	result, index, value = binarySearch(arr, 30)
	fmt.Printf("%t Index: %d, Value: %d\n", result, index, value)
}
