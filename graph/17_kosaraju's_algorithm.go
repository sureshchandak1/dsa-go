package graph

import (
	"container/list"
	"fmt"
)

// Kosaraju's Algorithm for Strongly Connected Components
// Directed Graph

func kosarajuAlgo(totalNode int, edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
	}

	// topological sort
	topoStack := list.New()
	visited := make(map[int]bool)

	for node := 0; node < totalNode; node++ {
		if !visited[node] {
			kosarajuAlgoTopoDFS(node, adj, visited, topoStack)
		}
	}

	// create a transpose graph
	transpose := make(map[int][]int)

	for node := 0; node < totalNode; node++ {
		visited[node] = false

		for _, nbr := range adj[node] {
			transpose[nbr] = append(transpose[nbr], node)
		}
	}

	// dfs call using above ordering
	count := 0
	for topoStack.Len() != 0 {

		pop := topoStack.Back()
		topoStack.Remove(pop)

		if !visited[pop.Value.(int)] {
			count++
			kosarajuAlgoDFS(pop.Value.(int), transpose, visited)
		}
	}

	fmt.Println(count)
}

func kosarajuAlgoTopoDFS(node int, adj map[int][]int, visited map[int]bool, topoStack *list.List) {

	visited[node] = true

	for _, nbr := range adj[node] {

		if !visited[nbr] {
			kosarajuAlgoTopoDFS(nbr, adj, visited, topoStack)
		}
	}

	topoStack.PushBack(node)
}

func kosarajuAlgoDFS(node int, adj map[int][]int, visited map[int]bool) {

	visited[node] = true

	for _, nbr := range adj[node] {

		if !visited[nbr] {
			kosarajuAlgoDFS(nbr, adj, visited)
		}
	}
}
