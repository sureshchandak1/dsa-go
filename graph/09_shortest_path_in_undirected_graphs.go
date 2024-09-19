package graph

import (
	"container/list"
	"fmt"
)

func shortestPathUndirectedGraph(edges [][]int, source, destination int) {

	adj = make(map[int][]int)

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// Start BFS

	visited := make(map[int]bool)
	parent := make(map[int]int)

	queue := list.New()

	queue.PushBack(source)
	visited[source] = true
	parent[source] = -1

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		for _, neighbour := range adj[front.Value.(int)] {
			// If neighbour not visited
			if !visited[neighbour] {
				queue.PushBack(neighbour)
				visited[neighbour] = true
				parent[neighbour] = front.Value.(int)
			}
		}
	}

	// prepare shortest path
	pathResult := []int{}

	currentNode := destination
	pathResult = append(pathResult, currentNode)

	for currentNode != source {
		currentNode = parent[currentNode]
		pathResult = append(pathResult, currentNode)
	}

	reverseArray(pathResult)

	fmt.Println(pathResult)
}

func reverseArray(arr []int) {

	for start, end := 0, len(arr)-1; start < end; start, end = start+1, end-1 {
		arr[start], arr[end] = arr[end], arr[start]
	}

}
