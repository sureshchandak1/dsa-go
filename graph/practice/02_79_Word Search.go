package practice

func wordExist(board [][]byte, word string) bool {

	for r := range board {
		for c := range board[r] {
			if board[r][c] == word[0] {
				if wordExistDFS(board, r, c, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func wordExistDFS(board [][]byte, r, c int, word string, wordIndex int) bool {

	// Base case
	if wordIndex == len(word) {
		return true // valid case
	}
	// out of bound cases
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return false
	}
	// Invalid cases
	if board[r][c] == ' ' || board[r][c] != word[wordIndex] {
		return false
	}

	ch := board[r][c]
	board[r][c] = ' ' // mark as visited

	// dfs calls
	up := wordExistDFS(board, r-1, c, word, wordIndex+1)
	right := wordExistDFS(board, r, c+1, word, wordIndex+1)
	down := wordExistDFS(board, r+1, c, word, wordIndex+1)
	left := wordExistDFS(board, r, c-1, word, wordIndex+1)

	if up || right || down || left {
		return true
	}

	board[r][c] = ch // backtracking

	return false
}
