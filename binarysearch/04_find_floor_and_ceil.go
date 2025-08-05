package binarysearch

/**

- floor: greates value equal to or smaller than target
- ceil: smallest value equal to or greater than target

**/

func findFloor(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	ans := -1
	for start <= end {

		if arr[mid] == target {
			return arr[mid]
		} else if arr[mid] < target {
			ans = arr[mid]
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func findCeil(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	ans := -1
	for start <= end {

		if arr[mid] == target {
			return arr[mid]
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			ans = arr[mid]
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

/**
https://leetcode.com/problems/find-smallest-letter-greater-than-target
**/
func findCeilLetter(arr []byte, target byte) byte {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	var ans byte = '$'
	for start <= end {

		if arr[mid] <= target {
			start = mid + 1
		} else {
			ans = arr[mid]
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	if ans == '$' {
		return arr[0]
	}

	return ans
}
