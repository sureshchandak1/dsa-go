package graph

import (
	"container/list"
	"fmt"
)

func kahnsAlgoTopoSort(totalNode int, edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
	}

	inDegrees := make([]int, totalNode) // array index represent node
	// find all in degrees
	for _, value := range adj {
		for _, node := range value {
			inDegrees[node]++
		}
	}

	queue := list.New()

	// push all 0 indegrees in queue
	for index, d := range inDegrees {
		if d == 0 {
			queue.PushBack(index)
		}
	}

	// start BFS

	ans := []int{}

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		ans = append(ans, front.Value.(int))

		// neighbour in degrees update
		for _, neighbour := range adj[front.Value.(int)] {

			inDegrees[neighbour]--

			if inDegrees[neighbour] == 0 {
				queue.PushBack(neighbour)
			}
		}
	}

	fmt.Println(ans)

}
