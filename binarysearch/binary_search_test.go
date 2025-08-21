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

func TestRotatedSortedArray(t *testing.T) {
	fmt.Println(minClockWiseRotated([]int{43, 2, 4, 5, 7, 11}))
	fmt.Println(maxClockWiseRotated([]int{43, 2, 4, 5, 7, 11}))
	fmt.Println(maxClockWiseRotated([]int{2, 4, 5, 7, 11, 43}))
	fmt.Println(minClockWiseRotatedDuplicates([]int{2, 2, 2, 0, 1}))
	fmt.Println(rotationCountInClockWise([]int{43, 2, 4, 5, 7, 11}))
	fmt.Println(rotationCountInAntiClockWise([]int{43, 2, 4, 5, 7, 11}))
	fmt.Println(searchRotatedArray([]int{43, 2, 4, 5, 7, 11}, 11))
	fmt.Println(searchRotatedArray([]int{43, 2, 4, 5, 7, 11}, 4))
	fmt.Println(searchRotatedArray([]int{43, 2, 4, 5, 7, 11}, 43))
}

func TestBookAllocation(t *testing.T) {
	fmt.Println(bookAllocation([]int{12, 34, 67, 90}, 2))
	fmt.Println(bookAllocation([]int{15, 17, 20}, 5))
}

func TestShipWithinDays(t *testing.T) {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
}

func TestMinKokoEatingSpeed(t *testing.T) {
	fmt.Println(minKokoEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minKokoEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
}

func TestSmallestDivisor(t *testing.T) {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{44, 22, 33, 11, 1}, 5))
}

func TestMinimizedMaximumDistributed(t *testing.T) {
	fmt.Println(minimizedMaximumDistributed([]int{11, 6}, 6))
	fmt.Println(minimizedMaximumDistributed([]int{15, 10, 10}, 7))
	fmt.Println(minimizedMaximumDistributed([]int{100000}, 1))
}
