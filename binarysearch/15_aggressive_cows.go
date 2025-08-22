package binarysearch

import "sort"

func aggressiveCows(stalls []int, cows int) int {

	n := len(stalls)

	if n < cows {
		return -1
	}

	sort.Ints(stalls)

	res := -1

	start := 1
	end := stalls[n-1] - stalls[0]

	mid := start + (end-start)/2

	for start <= end {

		if isCowsAllocationPossible(stalls, cows, mid) {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return res
}

func isCowsAllocationPossible(stalls []int, cows, minDistance int) bool {

	cowsCount := 1
	lastCowDistance := stalls[0]

	for i := 1; i < len(stalls); i++ {

		if stalls[i]-lastCowDistance >= minDistance {
			cowsCount += 1
			lastCowDistance = stalls[i]
		}

		if cowsCount >= cows {
			return true
		}
	}

	return false

}
