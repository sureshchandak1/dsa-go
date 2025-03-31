package hashmap

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	fmt.Println(mostFrequentElement([]int{2, 1, 7, 8, 2, 1, 9, 10, 9, 9, 2, 8, 2}))
}

func TestIsSubset(t *testing.T) {
	fmt.Println(isSubset([]int{1, 2, 7, 1, 6, 4, 3}, []int{1, 2, 1, 4, 1}))
	fmt.Println(isSubset([]int{1, 2, 7, 1, 6, 4, 3}, []int{1, 2, 1, 4}))
}

func TestPairCount(t *testing.T) {
	fmt.Println(pairCount1([]int{1, 5, 7, 1}, 6))
	fmt.Println(pairCount1([]int{2, 8, 5}, 10))
	fmt.Println()
	fmt.Println(pairCount2([]int{1, 5, 7, 1}, 6))
	fmt.Println(pairCount2([]int{2, 8, 5}, 10))
}

func TestFindMinNumberOfSets(t *testing.T) {
	fmt.Println(findMinNumberOfSets1([]int{2, 1, 4, 1, 6, 5, 5, 5}))
	fmt.Println(findMinNumberOfSets2([]int{2, 1, 4, 1, 6, 5, 5, 5}))
}

func TestCheckSumSubarray(t *testing.T) {
	fmt.Print(checkSumSubarray([]int{2, 8, 2, 6, -6, 3, 2}, 5))
}

func TestFindKSumSubarray(t *testing.T) {
	fmt.Print(findKSumSubarray([]int{2, 8, 2, 6, -6, 3, 2}, 5))
}

func TestFindLongestKSumSubarray(t *testing.T) {
	fmt.Print(findLongestKSumSubarray([]int{2, 8, 2, 6, -6, 3, 2}, 5))
}

func TestFindKSumSubarrayCount(t *testing.T) {
	fmt.Print(findKSumSubarrayCount([]int{2, 8, 2, 6, -6, 3, 2}, 5))
}
