package search

func peakIndexInMountainArray(arr []int) int {

	start := 0
	end := len(arr) - 1
	mid := start + (end-start)/2

	for start < end {

		if arr[mid] < arr[mid+1] {
			start = mid + 1 // if mid value small then next value it means peak value present in right side
		} else {
			end = mid
		}

		mid = start + (end-start)/2
	}

	return start
}
