package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMaxSumSubarray(t *testing.T) {
	fmt.Println(maxSumSubarray([]int{100, 200, 300, 400}, 2))
	fmt.Println(maxSumSubarray([]int{1, 5, 4, 2, 9, 9, 9}, 3))
}

func TestMinimumRecolors(t *testing.T) {
	fmt.Println(minimumRecolors("WBBWWBBWBW", 7))
	fmt.Println(minimumRecolors("WBWBBBW", 2))
}

func TestMaximumSubarraySum(t *testing.T) {
	fmt.Println(maximumSubarraySum([]int{4, 4, 4}, 3))
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
}
