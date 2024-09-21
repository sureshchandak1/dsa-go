package graph

import "fmt"

func articulationPointsAlgo(totalNode int, edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	result := []int{}

	visited := make(map[int]bool)
	disc := make([]int, totalNode) // discovery
	low := make([]int, totalNode)
	parent := -1
	timer := 0

	for i := 0; i < totalNode; i++ {
		disc[i] = -1
		low[i] = -1
	}

	// DFS

	for node := 0; node < totalNode; node++ {

		if !visited[node] {
			articulationDFS(node, parent, &timer, adj, visited, disc, low, &result)
		}
	}

	fmt.Println("Articulation Points: ", result)
}

func articulationDFS(node, parent int, timer *int, adj map[int][]int, visited map[int]bool, disc, low []int, result *[]int) {

	visited[node] = true
	*timer++
	disc[node] = *timer
	low[node] = *timer

	child := 0

	for _, nbr := range adj[node] {

		if nbr == parent {
			continue
		}

		if !visited[nbr] {

			articulationDFS(nbr, node, timer, adj, visited, disc, low, result)

			low[node] = min(low[node], low[nbr])

			// check edge is bridge

			if low[nbr] >= disc[node] && parent != -1 {
				*result = append(*result, node)
			}

			child++

		} else {
			// node already visited and not parent
			// Back edge
			low[node] = min(low[node], low[nbr])
		}
	}

	if parent == -1 && child > 1 {
		*result = append(*result, node)
	}
}
