package binarysearch

import "math"

func bookAllocation(bookPages []int, students int) int {

	if len(bookPages) < students {
		return -1
	}

	/**
	Find Range:
	Start = max of array
	End = Sum of array
	**/

	start := math.MinInt
	end := 0

	for i := 0; i < len(bookPages); i++ {
		if bookPages[i] > start {
			start = bookPages[i]
		}

		end = end + bookPages[i]
	}

	res := -1

	mid := start + (end-start)/2 // mid: max pages that can be allocated to single student.

	for start <= end {

		// If allocation is possible then minimize the number of pages
		if ifAllocationPossible(bookPages, mid, students) {
			res = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return res

}

func ifAllocationPossible(bookPages []int, maxPages, students int) bool {

	currentStudent := 1
	pages := 0

	for i := 0; i < len(bookPages); i++ {

		pages += bookPages[i]

		// If pages exceed maxPages
		if pages > maxPages {
			// allocate to next student
			currentStudent++
			pages = bookPages[i]
		}

		if currentStudent > students {
			return false
		}

	}

	return true

}
