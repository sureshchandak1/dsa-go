package practice

import (
	"fmt"
	"slices"
)

func findRelativeRanks(score []int) []string {

	n := len(score)
	maxVal := slices.Max(score)

	mapIndex := make([]int, maxVal+1)

	// fill index into map
	for i := 0; i < n; i++ {
		mapIndex[score[i]] = i + 1 // +1 to handle 0 index, array element default value also 0
	}

	result := make([]string, n)

	rank := 1
	for i := maxVal; i >= 0; i-- {
		if mapIndex[i] != 0 {
			originalIndex := mapIndex[i] - 1 // -1 to get original index
			if rank == 1 {
				result[originalIndex] = "Gold Medal"
			} else if rank == 2 {
				result[originalIndex] = "Silver Medal"
			} else if rank == 3 {
				result[originalIndex] = "Bronze Medal"
			} else {
				result[originalIndex] = fmt.Sprintf("%d", rank)
			}

			rank++
		}

		if rank > n {
			break
		}
	}

	return result
}
