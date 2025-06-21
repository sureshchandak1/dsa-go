package dp

import "slices"

func maxSquare(mat [][]int) int {
	// n := len(mat)
	// m := len(mat[0])

	maxi := 0
	// maxSquareRes(mat, 0, 0, &maxi)

	// dp := make([][]int, n+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int{-1}, m+1)
	// }
	// maxSquareDp(mat, 0, 0, &maxi, dp)

	// maxSquareTab(mat, &maxi)

	maxSquareOptimize(mat, &maxi)

	return maxi
}

func maxSquareRes(mat [][]int, row, col int, maxi *int) int {

	// Base case
	if row >= len(mat) || col >= len(mat[0]) {
		return 0
	}

	right := maxSquareRes(mat, row, col+1, maxi)
	diagonal := maxSquareRes(mat, row+1, col+1, maxi)
	down := maxSquareRes(mat, row+1, col, maxi)

	if mat[row][col] == 1 {
		ans := 1 + min(right, min(diagonal, down))
		*maxi = max(*maxi, ans)
		return ans
	} else {
		return 0
	}

}

func maxSquareDp(mat [][]int, row, col int, maxi *int, dp [][]int) int {

	// Base case
	if row >= len(mat) || col >= len(mat[0]) {
		return 0
	}

	if dp[row][col] != -1 {
		return dp[row][col]
	}

	right := maxSquareDp(mat, row, col+1, maxi, dp)
	diagonal := maxSquareDp(mat, row+1, col+1, maxi, dp)
	down := maxSquareDp(mat, row+1, col, maxi, dp)

	if mat[row][col] == 1 {
		dp[row][col] = 1 + min(right, min(diagonal, down))
		*maxi = max(*maxi, dp[row][col])
		return dp[row][col]
	} else {
		dp[row][col] = 0
		return dp[row][col]
	}

}

func maxSquareTab(mat [][]int, maxi *int) int {
	n := len(mat)
	m := len(mat[0])

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int{0}, m+1)
	}

	// according base case
	dp[n][m] = 0

	for row := n - 1; row >= 0; row-- {
		for col := m - 1; col >= 0; col-- {

			right := dp[row][col+1]
			diagonal := dp[row+1][col+1]
			down := dp[row+1][col]

			if mat[row][col] == 1 {
				dp[row][col] = 1 + min(right, min(diagonal, down))
				*maxi = max(*maxi, dp[row][col])
			} else {
				dp[row][col] = 0
			}

		}
	}

	return -1
}

func maxSquareOptimize(mat [][]int, maxi *int) int {

	n := len(mat)
	m := len(mat[0])

	curr := make([]int, m+1)
	next := make([]int, m+1)

	for row := n - 1; row >= 0; row-- {
		for col := m - 1; col >= 0; col-- {

			right := curr[col+1]
			diagonal := next[col+1]
			down := next[col]

			if mat[row][col] == 1 {
				curr[col] = 1 + min(right, min(diagonal, down))
				*maxi = max(*maxi, curr[col])
			} else {
				curr[col] = 0
			}

		}

		next = slices.Clone(curr)
	}

	return -1
}
