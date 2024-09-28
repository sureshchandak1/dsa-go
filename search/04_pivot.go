package search

func pivot(arr []int) (int, int) {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	for start < end {

		if arr[mid] >= arr[0] {
			start = mid + 1
		} else {
			end = mid
		}

		mid = start + (end-start)/2
	}

	return start, arr[start]

}
