package graph

import (
	"container/list"
	"fmt"
)

// Shortest Path in Directed Acyclic Graphs (DAG) with weight

type Pair struct {
	v      int
	weight int
}

func shortestPathDAG(totalNode int, edges [][]int, source int) {

	adjList := make(map[int][]Pair)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		adjList[u] = append(adjList[u], Pair{v: v, weight: w})
	}

	topoStack := list.New()

	topoSortDAG(totalNode, adjList, topoStack)

	path := getShortestPathDAG(totalNode, source, adjList, topoStack)

	fmt.Println(path)
}

func getShortestPathDAG(totalNode int, source int, adj map[int][]Pair, topo *list.List) []int {

	dist := make([]int, totalNode)

	for i := 0; i < totalNode; i++ {
		dist[i] = 2147483647
	}

	dist[source] = 0

	for topo.Len() != 0 {

		pop := topo.Back()
		topo.Remove(pop)

		if dist[pop.Value.(int)] != 2147483647 {

			for _, neighbour := range adj[pop.Value.(int)] {

				if dist[pop.Value.(int)]+neighbour.weight < dist[neighbour.v] {
					dist[neighbour.v] = dist[pop.Value.(int)] + neighbour.weight
				}
			}
		}

	}

	return dist
}

func topoSortDAG(totalNode int, adj map[int][]Pair, stack *list.List) {

	visited := make(map[int]bool)

	for node := 0; node < totalNode; node++ {
		if !visited[node] {
			topoSortDAG_DFS(node, adj, stack, visited)
		}
	}
}

func topoSortDAG_DFS(node int, adj map[int][]Pair, stack *list.List, visited map[int]bool) {

	visited[node] = true

	for _, neighbour := range adj[node] {

		if !visited[neighbour.v] {
			topoSortDAG_DFS(neighbour.v, adj, stack, visited)
		}
	}

	stack.PushBack(node)
}
