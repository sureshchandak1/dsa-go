package binarysearch

import "slices"

func minKokoEatingSpeed(piles []int, h int) int {

	start := 1
	end := slices.Max(piles)

	mid := start + (end-start)/2

	res := -1

	for start <= end {

		if isEatingSpeedValid(piles, h, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2

	}

	return res
}

func isEatingSpeedValid(piles []int, hours, eatingSpeed int) bool {

	hoursSpent := 0

	for i := 0; i < len(piles); i++ {

		hoursSpent = hoursSpent + piles[i]/eatingSpeed

		if piles[i]%eatingSpeed != 0 {
			hoursSpent += 1
		}

		if hoursSpent > hours {
			return false
		}

	}

	return true
}
