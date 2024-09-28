package search

// Binary search apply only on monotonic function
// element should be in monotonic function increase/decrease order

func binarySearch(arr []int, target int) (bool, int, int) {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	for start <= end {

		if arr[mid] == target {
			return true, mid, target
		}

		if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return false, -1, -1

}
