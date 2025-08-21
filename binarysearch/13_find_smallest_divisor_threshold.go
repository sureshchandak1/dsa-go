package binarysearch

import (
	"slices"
)

func smallestDivisor(nums []int, threshold int) int {

	start := 1
	end := slices.Max(nums)

	res := -1

	mid := start + (end-start)/2

	for start <= end {

		if isDivisionPossible(nums, threshold, mid) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return res
}

func isDivisionPossible(nums []int, threshold, divisor int) bool {

	sumOfDivision := 0

	for i := 0; i < len(nums); i++ {

		sumOfDivision = sumOfDivision + nums[i]/divisor

		if nums[i]%divisor != 0 {
			sumOfDivision += 1
		}

		if sumOfDivision > threshold {
			return false
		}
	}

	return true
}
