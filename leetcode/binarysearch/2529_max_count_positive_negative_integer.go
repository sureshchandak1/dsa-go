package binarysearch

func maximumCount(nums []int) int {
	size := len(nums)
	end := size - 1

	nPos := binarySearch(nums, 0, end, -1)
	pPos := binarySearch(nums, nPos, end, 0)

	pCount := size - pPos

	return max(nPos, pCount)
}

func binarySearch(nums []int, start, end, target int) int {

	var mid int
	for start <= end {
		mid = start + (end-start)/2

		if nums[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}

func maximumCountBruteForce(nums []int) int {
	size := len(nums)

	negCount := 0
	zeroCount := 0

	for index := 0; index < size; index++ {
		if nums[index] == 0 {
			zeroCount++
			continue
		}
		if nums[index] < 0 {
			negCount++
		} else {
			break
		}
	}

	posCount := size - negCount - zeroCount

	return max(negCount, posCount)
}
