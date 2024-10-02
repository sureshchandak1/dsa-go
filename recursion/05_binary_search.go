package recursion

/*
	return = index
*/
func binarySearch(arr []int, start, end, target int) int {

	// Base case
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if arr[mid] == target {
		return mid
	} else {
		if target > arr[mid] {
			rightPartSearch := binarySearch(arr, mid+1, end, target)
			return rightPartSearch
		} else {
			leftPartSearch := binarySearch(arr, start, mid-1, target)
			return leftPartSearch
		}
	}

}

func linearSearch(arr []int, target, index int) int {

	// Base case
	if index == len(arr) {
		return -1
	}

	if arr[index] == target {
		return index
	} else {
		remainingSearch := linearSearch(arr, target, index+1)
		return remainingSearch
	}
}

func isSorted(arr []int, index, size int) bool {

	// Base case
	if size == 0 || size == 1 {
		return true
	}

	if arr[index] > arr[index+1] {
		return false
	} else {
		remainingPart := isSorted(arr, index+1, size-1)
		return remainingPart
	}
}
