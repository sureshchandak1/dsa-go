package search

func searchRotatedArray(arr []int, target int) bool {

	size := len(arr)
	pIndex, _ := pivot(arr)

	if target >= arr[pIndex] && target <= arr[size-1] {
		return searchElement(arr, target, pIndex, size-1)
	} else {
		return searchElement(arr, target, 0, pIndex-1)
	}
}

func searchElement(arr []int, target, start, end int) bool {

	mid := start + (end-start)/2

	for start <= end {

		if arr[mid] == target {
			return true
		}

		if target > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return false
}
