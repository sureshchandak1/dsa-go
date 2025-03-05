package leetcode

func coloredCells(n int) int64 {
	return optimizedColoredCells(n)
}

func solveColoredCells(n int) int64 {
	if n == 1 {
		return 1
	}

	n--
	var x int64 = 4
	var ans int64 = 1

	for n > 0 {
		ans = ans + x
		x += 4
		n--
	}

	return ans
}

/*
1 + (4 * 1) + (4 * 2) + (4 * 3) ....
*/
func optimizedColoredCells(n int) int64 {
	return int64(1 + (n-1)*n*2)
}
