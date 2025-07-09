package dp

import (
	"slices"
	"sort"
)

func maxCuboidsHeight(cuboids [][]int) int {

	// step1: sort all dimenstions for every cuboid
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}

	// step2: sort all cuboids basis on w or l
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		} else if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	// step3: use LIS solution
	return maxCuboidsHeightLIS(cuboids)
}

func maxCuboidsHeightLIS(cuboids [][]int) int {

	check := func(base, newBox []int) bool {
		if newBox[0] <= base[0] && newBox[1] <= base[1] && newBox[2] <= base[2] {
			return true
		}
		return false
	}

	n := len(cuboids)

	currRow := make([]int, n+1)
	nextRow := make([]int, n+1)

	for curr := n - 1; curr >= 0; curr-- {

		for prev := curr - 1; prev >= -1; prev-- {

			include := 0
			if prev == -1 || check(cuboids[curr], cuboids[prev]) {
				include = cuboids[curr][2] + nextRow[curr+1]
			}

			exclude := nextRow[prev+1]

			currRow[prev+1] = max(include, exclude)

		}

		nextRow = slices.Clone(currRow)
	}

	return nextRow[0]

}
