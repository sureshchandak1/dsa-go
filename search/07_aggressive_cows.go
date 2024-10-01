package search

import "sort"

func aggressiveCows(stalls []int, totalCows int) int {

	sort.Slice(stalls, func(i, j int) bool {
		return stalls[i] < stalls[j]
	})

	maxi := stalls[0]
	for _, stall := range stalls {
		maxi = max(maxi, stall)
	}

	start := 0
	end := maxi
	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if isCowAllow(stalls, totalCows, mid) {
			ans = mid
			start = mid + 1 // move right, required largest distance
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans

}

func isCowAllow(stalls []int, totalCows, mid int) bool {

	cowCount := 1
	lastPosition := stalls[0]

	for _, stall := range stalls {

		if stall-lastPosition >= mid {
			cowCount++
			if cowCount == totalCows {
				return true
			}
			lastPosition = stall
		}
	}

	return false

}
