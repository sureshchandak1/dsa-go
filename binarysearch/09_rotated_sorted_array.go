package binarysearch

func minClockWiseRotated(arr []int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start < end {

		if mid > 0 && arr[mid] < arr[mid-1] {
			return mid
		} else if arr[end] > arr[mid] {
			// If right half is sorted
			end = mid - 1 // move left
		} else {
			start = mid + 1 // move right
		}

		mid = start + (end-start)/2
	}

	return start
}

func minClockWiseRotatedDuplicates(arr []int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start < end {

		if arr[start] == arr[mid] && arr[end] == arr[mid] {
			start++
			end--
			continue
		}

		if mid > 0 && arr[mid] < arr[mid-1] {
			return mid
		} else if arr[end] >= arr[mid] {
			// If right half is sorted
			end = mid - 1 // move left
		} else {
			start = mid + 1 // move right
		}

		mid = start + (end-start)/2
	}

	return start
}

func rotationCountInClockWise(arr []int) int {
	return minClockWiseRotated(arr)
}

func rotationCountInAntiClockWise(arr []int) int {
	index := minClockWiseRotated(arr)
	return (len(arr) - index) % len(arr)
}

func maxClockWiseRotated(arr []int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start < end {

		if mid < len(arr)-1 && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[start] > arr[mid] {
			// If right half is sorted
			end = mid - 1 // move left
		} else {
			start = mid + 1 // move right
		}

		mid = start + (end-start)/2
	}

	return end
}

func searchRotatedArray(arr []int, target int) int {

	start := 0
	end := len(arr) - 1

	mid := start + (end-start)/2

	for start <= end {

		if arr[mid] == target {
			return mid
		}

		// left side is sorted
		if arr[start] <= arr[mid] {
			// can ans be found in left side.
			if target >= arr[start] && target < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			// right side is sorted
			// can ans be found in right side
			if target > arr[mid] && target <= arr[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		mid = start + (end-start)/2
	}

	return -1
}
