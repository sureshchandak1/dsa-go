package recursion

var (
	direction = "DLRU"
	dr        = []int{1, 0, 0, -1}
	dc        = []int{0, -1, 1, 0}
)

func findRatPath(board [][]int, size int) []string {

	result := []string{}

	if board[0][0] != 0 && board[size-1][size-1] != 0 {

		row := 0
		col := 0
		path := ""

		findPath(board, size, row, col, path, &result)
	}

	return result
}

func findPath(board [][]int, size, row, col int, path string, result *[]string) {

	// Base case
	if row == size-1 && col == size-1 {
		*result = append(*result, string(path))
		return
	}

	// Mark the current cell as blocked
	board[row][col] = 0

	for i := 0; i < 4; i++ {

		nextRow := row + dr[i]
		nextCol := col + dc[i]

		if isSafe(board, size, nextRow, nextCol) {

			path = path + string(direction[i])

			findPath(board, size, nextRow, nextCol, path, result)

			// Remove the last direction when backtracking
			path = path[:len(path)-1]
		}
	}

	// Mark the current cell as unblocked
	board[row][col] = 1

}

func isSafe(board [][]int, size, row, col int) bool {
	return (row >= 0 && row < size) && (col >= 0 && col < size) && board[row][col] == 1
}
