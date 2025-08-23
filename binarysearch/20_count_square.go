package binarysearch

func countSquare(n int) int {

	sqrt := func(x int) int { // calculate square root

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

		return start // ceil

	}(n)

	return sqrt - 1
}
