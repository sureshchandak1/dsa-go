package practice

import "sort"

func maximumImportance(n int, roads [][]int) int64 {

	// find the degree
	degree := make([]int64, n)

	for _, road := range roads {
		u := road[0]
		v := road[1]

		degree[u]++
		degree[v]++
	}

	sort.Slice(degree, func(i, j int) bool {
		return degree[i] < degree[j]
	})

	var result int64 = 0
	var label int64 = 1
	for i := range n {
		result += degree[i] * label
		label++
	}

	return result
}
