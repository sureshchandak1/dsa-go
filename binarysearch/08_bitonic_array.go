package binarysearch

func bitonicArray(arr []int) int {
	return -1
}

// Bitonic point
func peakInBitonicArray(arr []int) int {

	n := len(arr)
	start := 0
	end := n - 1

	mid := start + (end-start)/2

	for start <= end {

		next := (mid + 1) % n
		prev := (mid - 1 + n) % n

		if arr[mid] > arr[prev] && arr[mid] > arr[next] {
			return mid
		} else if arr[mid] > arr[prev] {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return -1

}

func binarySearchBitonicArray(arr []int, target int) int {

	peak := peakInBitonicArray(arr)

	if peak == -1 || target > arr[peak] {
		return -1
	}

	if target == arr[peak] {
		return peak
	}

	var ans int = -1
	ans = binarySearchInsc(arr, target, 0, peak-1) // search in left half

	if ans == -1 {
		ans = binarySearchDesc(arr, target, peak+1, len(arr)-1) // search in right half
	}

	return ans
}

func minBitonicArray(arr []int) int {

	n := len(arr)

	if n == 0 {
		return -1
	}
	if n == 1 {
		return arr[0]
	}

	return min(arr[0], arr[n-1])
}

func binarySearchInsc(arr []int, target, start, end int) int {

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

func binarySearchDesc(arr []int, target, start, end int) int {

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
