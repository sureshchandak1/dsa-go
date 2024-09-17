package graph

import (
	"container/list"
	"fmt"
)

// Directed Acyclic Graph (DAG), Acyclic = no cycle
// Linear ordering of vertices such that for every edge u -> v, u always appears before v in that ordering
// Topological sort only apply on DAG graph
func topologicalSortDFS(totalNodes int, edges [][]int) {

	// Create adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
	}

	visited := make(map[int]bool)
	stack := list.New()

	for node := 0; node < totalNodes; node++ {

		if !visited[node] {
			topoSort(node, adj, stack, visited)
		}
	}

	sortResult := []int{}

	for stack.Len() != 0 {
		pop := stack.Back()
		sortResult = append(sortResult, pop.Value.(int))
		stack.Remove(pop)
	}

	fmt.Println(sortResult)

}

func topoSort(node int, adj map[int][]int, stack *list.List, visited map[int]bool) {

	visited[node] = true

	for _, neighbour := range adj[node] {

		if !visited[neighbour] {
			topoSort(neighbour, adj, stack, visited)
		}
	}

	stack.PushBack(node)
}
