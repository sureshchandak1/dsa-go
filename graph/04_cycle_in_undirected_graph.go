package graph

import (
	"container/list"
	"fmt"
)

func isCyclicBFSUndirectedGraph(totalNodes int, edges [][]int) {

	// adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool)

	for node := 0; node < totalNodes; node++ {
		if !visited[node] {
			isCycle := isCyclicBFS(node, adj, visited)
			if isCycle {
				fmt.Println("BFS result: Yes")
				return
			}
		}
	}

	fmt.Println("BFS result: No")

}

func isCyclicBFS(node int, adj map[int][]int, visited map[int]bool) bool {

	parent := make(map[int]int)
	queue := list.New()

	visited[node] = true
	parent[node] = -1
	queue.PushBack(node)

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		for _, neighbour := range adj[front.Value.(int)] {

			if visited[neighbour] && neighbour != parent[front.Value.(int)] {
				return true
			} else if !visited[neighbour] {
				queue.PushBack(neighbour)
				visited[neighbour] = true
				parent[neighbour] = front.Value.(int)
			}
		}
	}

	return false

}

func isCyclicDFSUndirectedGraph(totalNodes int, edges [][]int) {

	// adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visited := make(map[int]bool)

	for node := 0; node < totalNodes; node++ {
		if !visited[node] {
			isCycle := isCyclicDFS(node, -1, adj, visited)
			if isCycle {
				fmt.Println("DFS result: Yes")
				return
			}
		}
	}

	fmt.Println("DFS result: No")

}

func isCyclicDFS(node int, parent int, adj map[int][]int, visited map[int]bool) bool {

	visited[node] = true

	for _, neighbour := range adj[node] {

		if !visited[neighbour] {
			isCycle := isCyclicDFS(neighbour, node, adj, visited)
			if isCycle {
				return true
			}
		} else if neighbour != parent {
			return true
		}
	}

	return false
}
