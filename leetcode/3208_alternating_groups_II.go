package leetcode

func numberOfAlternatingGroups(colors []int, k int) int {

	result := 0
	size := len(colors)

	left := 0
	for right := 1; right < (size + k - 1); right++ {
		// skip subarray
		if colors[right%size] == colors[(right-1)%size] {
			left = right
		}

		if right-left+1 == k {
			result++
			left++ // move to next subarray or shrinking phase
		}
	}

	return result
}
