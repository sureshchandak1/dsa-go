package leetcode

func findMissingAndRepeatedValues(grid [][]int) []int {
	result := []int{}

	n := len(grid)
	sq := n * n
	var set = make([]int, sq+1)
	currSum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if set[grid[i][j]] != 0 {
				// repeated number
				result = append(result, grid[i][j])
			} else {
				set[grid[i][j]] = 1
				currSum += grid[i][j]
			}
		}
	}

	totalSum := sq * (sq + 1) / 2
	// missing number
	result = append(result, totalSum-currSum)

	return result
}
