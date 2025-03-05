package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func TestPartitionByPivot(t *testing.T) {
	arr := []int{9, 12, 5, 10, 14, 3, 10}
	fmt.Println(partitionByPivot1(arr, 10))
	fmt.Println(partitionByPivot2(arr, 10))

	arr = []int{-3, 4, 3, 2}
	fmt.Println(partitionByPivot1(arr, 2))
	fmt.Println(partitionByPivot2(arr, 2))
}
