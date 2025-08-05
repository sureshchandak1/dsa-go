package binarysearch

func firstIndex(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	ans := -1
	for start <= end {

		if arr[mid] == target {
			ans = mid
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func lastIndex(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	ans := -1
	for start <= end {

		if arr[mid] == target {
			ans = mid
			start = mid + 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}
