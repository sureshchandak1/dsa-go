package graph

import (
	"container/list"
	"fmt"
)

func cycleByKahnsAlgo(totalNode int, edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		adj[u] = append(adj[u], v)
	}

	inDegrees := make([]int, totalNode+1) // array index represent node
	// find all in degrees
	for _, value := range adj {
		for _, node := range value {
			inDegrees[node]++
		}
	}

	queue := list.New()

	// push all 0 indegrees in queue
	// index start from 1 because nodes start from 1
	for index := 1; index <= totalNode; index++ {
		if inDegrees[index] == 0 {
			queue.PushBack(index)
		}
	}

	// start BFS

	count := 0

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		count++

		// neighbour in degrees update
		for _, neighbour := range adj[front.Value.(int)] {

			inDegrees[neighbour]--

			if inDegrees[neighbour] == 0 {
				queue.PushBack(neighbour)
			}
		}
	}

	if count == totalNode {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}

func cycleByKahnsAlgo2(totalNode int, edges [][]int) {

	adj = make(map[int][]int)

	for _, edge := range edges {

		u := edge[0] - 1 // nodes starts from 1
		v := edge[1] - 1

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

	count := 0

	for queue.Len() != 0 {

		front := queue.Front()
		queue.Remove(front)

		count++

		// neighbour in degrees update
		for _, neighbour := range adj[front.Value.(int)] {

			inDegrees[neighbour]--

			if inDegrees[neighbour] == 0 {
				queue.PushBack(neighbour)
			}
		}
	}

	if count == totalNode {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}

}
