package binarysearch

func minCapability(nums []int, k int) int {
	size := len(nums)
	min, max := 2342342342, -2342342342
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	result := 0

	start, end, mid := min, max, 0

	for start <= end {

		mid = start + (end-start)/2

		if isRobberyPossible(nums, mid, k, size) {
			result = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result
}

func isRobberyPossible(nums []int, capability, minHouses, size int) bool {

	housesRobbed := 0

	for i := 0; i < size; i++ {
		if nums[i] <= capability {
			housesRobbed++
			i++ // skip for adjacent condition (ignore neaighbour house)
		}

		if housesRobbed >= minHouses {
			return true
		}
	}

	return false
}
