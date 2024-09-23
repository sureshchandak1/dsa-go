package graph

import "fmt"

func bellmanFordShortestPath(totalNode int, edges [][]int, source, destination int) {

	largeInt := 1e8
	dist := make([]int, totalNode)

	// Insert default value
	for i := 0; i < totalNode; i++ {
		dist[i] = int(largeInt)
	}

	dist[source-1] = 0 // source - 1 , node is start from 1

	// n - 1 times

	for i := 1; i <= totalNode; i++ {
		// traverse on edge list
		for j := 0; j < len(edges); j++ {

			u := edges[j][0] - 1 // -1, node is start from 1
			v := edges[j][1] - 1 // -1, node is start from 1
			w := edges[j][2]

			if dist[u] != int(largeInt) && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}

	// check for negative cycle

	flag := false
	for _, edge := range edges {

		u := edge[0] - 1
		v := edge[1] - 1
		w := edge[2]

		if dist[u] != int(largeInt) && dist[u]+w < dist[v] {
			flag = true
		}
	}

	if !flag {
		fmt.Println(dist)
		fmt.Println(dist[destination-1])
		return
	}

	fmt.Println(-1)

}
