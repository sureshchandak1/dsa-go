package binarysearch

import "math"

func minAbsoluteDiff1(arr []int, target int) int {

	if len(arr) == 0 {
		return 0
	}

	floor := findFloor(arr, target)
	ceil := findCeil(arr, target)

	if floor == -1 {
		return int(math.Abs(float64(target - ceil)))
	}
	if ceil == -1 {
		return int(math.Abs(float64(target - floor)))
	}

	return int(min(math.Abs(float64(target-ceil)), math.Abs(float64(target-floor))))

}

func minAbsoluteDiffOptimized(arr []int, target int) int {

	if len(arr) == 0 {
		return 0
	}

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start <= end {

		if arr[mid] == target {
			return int(math.Abs(float64(target - arr[mid])))
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	/**
	start - pointing to ceil index
	end - pointing to floor index
	**/

	if end == -1 {
		return int(math.Abs(float64(target - arr[start])))
	}
	if start == len(arr) {
		return int(math.Abs(float64(target - arr[end])))
	}

	return int(min(math.Abs(float64(target-arr[start])), math.Abs(float64(target-arr[end]))))
}
