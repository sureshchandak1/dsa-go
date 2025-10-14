package practice

func findStarGraphCenterOptimized(edges [][]int) int {

	pair1 := edges[0]
	pair2 := edges[1]

	if pair1[0] == pair2[0] || pair1[0] == pair2[1] {
		return pair1[0]
	}

	return pair1[1]
}

func findStarGraphCenter(edges [][]int) int {

	totalEdges := len(edges)

	freqMap := make(map[int]int)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		freqMap[u]++
		freqMap[v]++

		if freqMap[u] == totalEdges {
			return u
		}
		if freqMap[v] == totalEdges {
			return v
		}
	}

	return -1
}
