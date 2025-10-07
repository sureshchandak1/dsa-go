package practice

func islandPerimeter(grid [][]int) int {

	rows := len(grid)
	cols := len(grid[0])

	count := 0

	for r := range rows {
		for c := range cols {

			if grid[r][c] == 1 {

				// up
				if (r > 0 && grid[r-1][c] == 0) || r == 0 {
					count++
				}

				// right
				if (c < cols-1 && grid[r][c+1] == 0) || c == cols-1 {
					count++
				}

				// down
				if (r < rows-1 && grid[r+1][c] == 0) || r == rows-1 {
					count++
				}

				// left
				if (c > 0 && grid[r][c-1] == 0) || c == 0 {
					count++
				}
			}
		}
	}

	return count
}
