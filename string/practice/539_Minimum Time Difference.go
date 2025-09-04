package practice

import (
	"math"
	"strconv"
)

func findMinDifference(timePoints []string) int {

	mins := make([]bool, 1440)

	for _, time := range timePoints {
		h, _ := strconv.Atoi(time[:2])
		m, _ := strconv.Atoi(time[3:5])

		minutes := h*60 + m

		if mins[minutes] {
			return 0
		}

		mins[minutes] = true
	}

	prev := -1
	minDiff := math.MaxInt32
	firstVal := -1

	for curr := 0; curr < 1440; curr++ {
		if mins[curr] {
			if prev == -1 {
				firstVal = curr
				prev = curr
			} else {
				minDiff = min(minDiff, curr-prev)
				prev = curr
			}
		}
	}

	if prev != -1 {
		minDiff = min(minDiff, 1440+firstVal-prev)
	}

	return minDiff
}
