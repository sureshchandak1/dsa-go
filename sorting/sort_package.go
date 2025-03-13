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

func sort2dArray(matrix [][]int) {
	sort.Slice(matrix[:], func(i, j int) bool {
		for x := range matrix[i] {
			if matrix[i][x] == matrix[j][x] {
				continue
			}
			return matrix[i][x] < matrix[j][x]
		}
		return false
	})
}
