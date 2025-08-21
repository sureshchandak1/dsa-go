package binarysearch

import "slices"

func minimizedMaximumDistributed(quantities []int, n int) int {

	start := 1
	end := slices.Max(quantities)

	res := -1

	mid := start + (end-start)/2

	for start <= end {

		if isDistributionPossible(quantities, n, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return res
}

func isDistributionPossible(quantities []int, stores, maxProducts int) bool {

	storeCount := 0

	for i := 0; i < len(quantities); i++ {

		storeCount = storeCount + quantities[i]/maxProducts

		if quantities[i]%maxProducts != 0 {
			storeCount += 1
		}

		if storeCount > stores {
			return false
		}
	}

	return true
}
