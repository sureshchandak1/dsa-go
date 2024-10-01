package search

func painterPartition(boards []int, totalPainter int) int {

	start := 0
	end := 0
	for _, board := range boards {
		end = end + board
	}

	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if isPainterAllow(boards, totalPainter, mid) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func isPainterAllow(boards []int, totalPainter, mid int) bool {

	partitionSum := 0
	painterCount := 1

	for _, board := range boards {

		if partitionSum+board <= mid {
			partitionSum = partitionSum + board
		} else {
			painterCount++

			if painterCount > totalPainter || board > mid {
				return false
			}

			partitionSum = board
		}
	}

	return true
}
