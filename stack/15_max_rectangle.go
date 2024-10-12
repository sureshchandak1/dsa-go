package stack

func maximalRectangle(matrix [][]int) int {

	// area for first row
	area := largestRectangleArea(matrix[0]) // mathod call from file no. 13

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] != 0 {
				matrix[i][j] = matrix[i][j] + matrix[i-1][j]
			} else {
				matrix[i][j] = 0
			}
		}

		newArea := largestRectangleArea(matrix[i])

		area = max(area, newArea)
	}

	return area
}
