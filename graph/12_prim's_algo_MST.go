package graph

import "fmt"

// Prim's Algorithm Minimum Spanning Tree (MST)
// Use: find MST
// Spanning Tree: When you convert a graph into a tree such that it contain n nodes and n-1 edges
// Convert graph into tree, where every node rechable

func primsAlgoMST(totalNode int, edges [][]int) {

	adjList := make(map[int][]Pair)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		adjList[u] = append(adjList[u], Pair{v: v, weight: w})
		adjList[v] = append(adjList[v], Pair{v: u, weight: w})
	}

	key := make([]int, totalNode+1)
	mst := make([]bool, totalNode+1)
	parent := make([]int, totalNode+1)

	for i := 0; i <= totalNode; i++ {
		key[i] = 2147483647
		mst[i] = false
		parent[i] = -1
	}

	// Start Algo

	key[1] = 0
	parent[1] = -1

	for index := 1; index < totalNode; index++ {

		mini := 2147483647
		miniNode := 0
		// find the min node in key array
		for v := 1; v <= totalNode; v++ {
			if !mst[v] && key[v] < mini {
				mini = key[v]
				miniNode = v
			}
		}

		// mark min node as true
		mst[miniNode] = true

		// check its adjacent nodes
		for _, neighbour := range adjList[miniNode] {

			v := neighbour.v      // node
			w := neighbour.weight // weight

			if !mst[v] && w < key[v] {
				parent[v] = miniNode
				key[v] = w
			}
		}
	}

	result := [][]int{}

	for index := 2; index <= totalNode; index++ {
		list := []int{parent[index], index, key[index]}

		result = append(result, list)
	}

	fmt.Println(result)
}
