package graph

import "fmt"

func isCyclicDFSDirectedGraph(totalNodes int, edges [][]int) {

	// Create adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
	}

	visited := make(map[int]bool)
	dfsVisited := make(map[int]bool)

	for node := 0; node < totalNodes; node++ {

		if !visited[node] {
			isCycle := checkCycleDFS(node, adj, visited, dfsVisited)
			if isCycle {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")

}

func checkCycleDFS(node int, adj map[int][]int, visited map[int]bool, dfsVisited map[int]bool) bool {

	visited[node] = true
	dfsVisited[node] = true

	for _, neighbour := range adj[node] {

		if !visited[neighbour] {
			isCycle := checkCycleDFS(neighbour, adj, visited, dfsVisited)
			if isCycle {
				return true
			}
		} else if dfsVisited[neighbour] {
			return true
		}
	}

	return false
}
