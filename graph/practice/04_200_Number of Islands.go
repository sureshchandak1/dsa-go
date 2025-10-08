package practice

func numIslands(grid [][]byte) int {

	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] != '0' && grid[r][c] != '2' {
				dfsNumIslands(grid, r, c)
				count++
			}
		}
	}

	return count
}

func dfsNumIslands(grid [][]byte, row, col int) {

	// Base case
	// out of bound
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return
	}
	// water or visited
	if grid[row][col] == '0' || grid[row][col] == '2' {
		return
	}

	grid[row][col] = '2' // mark visited

	dfsNumIslands(grid, row-1, col) // Up

	dfsNumIslands(grid, row, col+1) // Right

	dfsNumIslands(grid, row+1, col) // Down

	dfsNumIslands(grid, row, col-1) // Left
}
