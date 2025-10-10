package practice

func validPath(n int, edges [][]int, source int, destination int) bool {

	// adjacency list
	adj := make([][]int, n)
	for i := range n {
		adj[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make([]bool, n)

	return dfsValidPath(adj, source, destination, visited)
}

func dfsValidPath(adj [][]int, curNode int, destination int, visited []bool) bool {

	visited[curNode] = true // Mark visited

	if curNode == destination {
		return true
	}

	// visit all the neighbours of curNode
	for _, neighbour := range adj[curNode] {
		if !visited[neighbour] {
			if dfsValidPath(adj, neighbour, destination, visited) {
				return true
			}
		}
	}

	return false
}
