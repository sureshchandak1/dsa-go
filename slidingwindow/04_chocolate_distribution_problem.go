package slidingwindow

import "sort"

func chocolateDistribution(arr []int, m int) int {
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	n := len(arr)
	result := 24234234

	for i := range n - m + 1 {
		// arr[i] = window min element
		// arr[i+m-1] = window max element
		result = min(result, arr[i+m-1]-arr[i])
	}

	return result
}
