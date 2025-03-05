package sorting

import "sort"

func sortIntArray(arr []int, isAscending bool) []int {
	if isAscending {
		sort.Ints(arr)
	} else {
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})
	}
	return arr
}
