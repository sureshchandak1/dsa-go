package binarysearch

func binarySearchInfiniteArr(arr []int, target int) int {

	rangeIndex := findRangeInfiniteArr(arr, target)

	ans := binarySearch(arr, target, rangeIndex[0], rangeIndex[1])

	return ans
}

func findRangeInfiniteArr(arr []int, target int) []int {

	startEnd := make([]int, 2)

	start := 0
	end := 1

	for arr[end] < target {

		start = end
		if 2*end < len(arr) {
			end = 2 * end
		} else {
			end = len(arr) - 1
			break
		}

	}

	startEnd[0] = start
	startEnd[1] = end

	return startEnd
}

func binarySearch(arr []int, target, start, end int) int {

	mid := start + (end-start)/2

	for start <= end {

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return -1
}
