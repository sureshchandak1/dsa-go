package graph

import "fmt"

// Depth First Search (DFS)
func dfsTraversal(totalNodes int, edges [][]int) {

	// adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := [][]int{}
	visited := make(map[int]bool)

	for node := 0; node < totalNodes; node++ {
		if !visited[node] {
			component := []int{}
			dfs(node, adj, visited, &component)
			ans = append(ans, component)
		}
	}

	fmt.Println(ans)

}

func dfs(node int, adj map[int][]int, visited map[int]bool, component *[]int) {

	// ans store
	*component = append(*component, node)
	visited[node] = true // mark visited

	// for connected node recursive call
	for _, neighbour := range adj[node] {
		if !visited[neighbour] {
			dfs(neighbour, adj, visited, component)
		}
	}
}
