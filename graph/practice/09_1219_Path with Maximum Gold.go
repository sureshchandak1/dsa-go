package practice

import "math"

var dir = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var rows = 0
var cols = 0

func getMaximumGold(grid [][]int) int {

	rows = len(grid)
	cols = len(grid[0])

	totalGold := 0
	for i := range rows {
		for j := range cols {
			totalGold += grid[i][j]
		}
	}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	maxGold := 0

	for i := range rows {
		for j := range cols {

			if grid[i][j] != 0 {
				maxGold = int(math.Max(float64(maxGold), float64(dfsMaximumGold(i, j, grid, visited))))

				if maxGold == totalGold {
					return maxGold
				}
			}
		}
	}

	return maxGold

}

func dfsMaximumGold(row, col int, grid [][]int, visited [][]bool) int {

	// Base case
	// out of bound, visited, 0 gold
	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 0 || visited[row][col] {
		return 0
	}

	visited[row][col] = true // Mark visited

	maxGold := 0
	for i := range 4 {
		nextRow := row + dir[i][0]
		nextCol := col + dir[i][1]

		maxGold = int(math.Max(float64(maxGold), float64(dfsMaximumGold(nextRow, nextCol, grid, visited))))
	}

	visited[row][col] = false // Backtracking

	return grid[row][col] + maxGold
}
