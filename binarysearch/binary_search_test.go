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

func TestFirstLastIndex(t *testing.T) {
	arr := []int{2, 2, 3, 4, 6, 8, 9}
	fmt.Println(firstIndex(arr, 2))
	fmt.Println(lastIndex(arr, 9))

	fmt.Println(firstIndex(arr, 3))
	fmt.Println(lastIndex(arr, 8))
}

func TestCountFrequency(t *testing.T) {
	arr := []int{2, 2, 3, 4, 6, 8, 9, 9, 9, 9}
	fmt.Println(countFrequency(arr, 2))
	fmt.Println(countFrequency(arr, 4))
	fmt.Println(countFrequency(arr, 9))
	fmt.Println(countFrequency(arr, 60))
}

func TestFloorAndCeil(t *testing.T) {
	arr := []int{2, 2, 3, 4, 6, 8, 9, 9, 9, 9}
	fmt.Println(findFloor(arr, 7))
	fmt.Println(findFloor(arr, 12))
	fmt.Println(findFloor(arr, 1))
	fmt.Println(findCeil(arr, 7))
	fmt.Println(findCeil(arr, 12))
	fmt.Println(findCeil(arr, 1))
	fmt.Println(findCeilLetter([]byte{'x', 'x', 'y', 'y'}, 'z'))
}
