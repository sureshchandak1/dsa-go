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

func TestMinimumAbsoluteDifference(t *testing.T) {
	arr := []int{2, 4, 8, 13, 15, 17, 19}
	fmt.Println(minAbsoluteDiff1(arr, 12))
	fmt.Println(minAbsoluteDiffOptimized(arr, 12))
	fmt.Println(minAbsoluteDiffOptimized(arr, 13))
	fmt.Println(minAbsoluteDiffOptimized(arr, 25))
	fmt.Println(minAbsoluteDiffOptimized(arr, 0))
}

func TestBinarySearchInfiniteArr(t *testing.T) {
	arr := []int{2, 2, 4, 5, 7, 9, 25, 56, 88, 94}
	fmt.Println(binarySearchInfiniteArr(arr, 25))
	fmt.Println(binarySearchInfiniteArr(arr, 100))
}

func TestFindFirstIndexOf1(t *testing.T) {
	fmt.Println(findFirstIndexOf1([]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1}))
	fmt.Println(findFirstIndexOf1([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}))
	fmt.Println(findFirstIndexOf1([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}

func TestBitonicArray(t *testing.T) {
	arr := []int{2, 4, 6, 8, 11, 16, 13, 11, 9, 3, 2, 1}
	fmt.Println(minBitonicArray(arr))
	fmt.Println(peakInBitonicArray(arr))
	fmt.Println(binarySearchBitonicArray(arr, 16))
	fmt.Println(binarySearchBitonicArray(arr, 11))
}
