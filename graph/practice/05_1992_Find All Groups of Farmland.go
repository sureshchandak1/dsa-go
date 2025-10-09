package practice

func findFarmland(land [][]int) [][]int {

	n := len(land)
	m := len(land[0])

	var result [][]int

	for i := range n {
		for j := range m {

			if land[i][j] == 1 {

				arr := make([]int, 4)

				// start
				arr[0] = i
				arr[1] = j

				r := i
				c := j

				// mark visited
				for r = i; r < n && land[r][j] == 1; r++ {
					for c = j; c < m && land[r][c] == 1; c++ {
						land[r][c] = 2
					}
				}

				// end
				arr[2] = r - 1
				arr[3] = c - 1

				result = append(result, arr)
			}
		}
	}

	return result
}
