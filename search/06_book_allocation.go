package search

func bookAllocation(bookPages []int, totalStudent int) int {

	// we need find search space

	start := 0
	end := 0
	for _, pages := range bookPages {
		end = end + pages
	}

	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if isAllocationPositble(bookPages, totalStudent, mid) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func isAllocationPositble(bookPages []int, totalStudent, mid int) bool {

	totalAllocation := 0
	studentCount := 1

	for _, pages := range bookPages {

		if totalAllocation+pages <= mid {
			totalAllocation = totalAllocation + pages
		} else {
			studentCount++
			if studentCount > totalStudent || pages > mid {
				return false
			}

			totalAllocation = pages
		}
	}

	return true

}
