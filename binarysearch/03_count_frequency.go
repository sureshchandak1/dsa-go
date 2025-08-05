package binarysearch

func countFrequency(arr []int, target int) int {

	firstIndex := firstLastIndex(arr, target, true)

	if firstIndex == -1 {
		return 0
	}

	lastIndex := firstLastIndex(arr, target, false)

	return lastIndex - firstIndex + 1
}

func firstLastIndex(arr []int, target int, isFirst bool) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	ans := -1
	for start <= end {

		if arr[mid] == target {
			ans = mid
			if isFirst {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}
