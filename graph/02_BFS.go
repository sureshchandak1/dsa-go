package graph

import (
	"container/list"
	"fmt"
)

// Breadth First Search (BFS)
func bfsTraversal(edges [][]int) {

	// adjacency list
	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := []int{}
	visited := make(map[int]bool)

	queue := list.New()

	queue.PushBack(0)
	visited[0] = true

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		// store front node into ans
		ans = append(ans, front.Value.(int))

		// traverse all neighbours of frontNode
		for _, node := range adj[front.Value.(int)] {
			if !visited[node] {
				queue.PushBack(node)
				visited[node] = true
			}
		}
	}

	fmt.Println(ans)

}
