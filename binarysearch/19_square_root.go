package binarysearch

func mySqrt(x int) int {

	start := 1
	end := x

	for start <= end {

		mid := start + (end-start)/2

		if mid <= x/mid {
			if x%mid == 0 && mid == x/mid {
				return mid
			}
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return end
}
