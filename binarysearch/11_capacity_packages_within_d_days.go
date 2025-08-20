package binarysearch

import "math"

func shipWithinDays(weights []int, days int) int {

	start := math.MinInt
	end := 0

	for i := range weights {
		if weights[i] > start {
			start = weights[i]
		}

		end = end + weights[i]
	}

	res := -1

	mid := start + (end-start)/2

	for start <= end {

		if ifShipmentPossible(weights, days, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return res
}

func ifShipmentPossible(weights []int, days int, maxCapacity int) bool {

	currentDay := 1
	capacity := 0

	for i := 0; i < len(weights); i++ {

		capacity += weights[i]

		if capacity > maxCapacity {
			currentDay += 1
			capacity = weights[i]
		}

		if currentDay > days {
			return false
		}
	}

	return true
}
