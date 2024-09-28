package search

func firstLastPosition(arr []int, element int) (int, int, int) {

	first := firstPosition(arr, element)
	last := lastPosition(arr, element)

	return first, last, element
}

func firstPosition(arr []int, element int) int {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	firstIndex := -1

	for start <= end {

		if arr[mid] == element {
			firstIndex = mid
			end = mid - 1
		} else if arr[mid] < element {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return firstIndex
}

func lastPosition(arr []int, element int) int {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	lastIndex := -1

	for start <= end {

		if arr[mid] == element {
			lastIndex = mid
			start = mid + 1
		} else if arr[mid] < element {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return lastIndex
}
