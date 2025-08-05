package binarysearch

func binarySearchOrderAgnostic(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	if arr[start] <= arr[end] {
		return binarySearchInc(arr, target)
	} else {
		return binarySearchDec(arr, target)
	}
}

func binarySearchInc(arr []int, target int) int {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if arr[mid] == target {
			ans = mid
			break
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func binarySearchDec(arr []int, target int) int {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if arr[mid] == target {
			ans = mid
			break
		} else if arr[mid] < target {
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return ans
}
