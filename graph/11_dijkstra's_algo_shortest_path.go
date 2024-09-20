package graph

import (
	"container/heap"
	"fmt"
)

// Shortest path using Dijkstras Algorithm
func dijkstraShortPathAlgo(totalNode int, edges [][]int, source int) {

	// Create adjacency list
	adjList := make(map[int][]Pair)

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]
		w := edge[2]

		adjList[u] = append(adjList[u], Pair{v: v, weight: w})
		adjList[v] = append(adjList[v], Pair{v: u, weight: w})
	}

	// creation of distance array with infinite value initially
	distance := make([]int, totalNode)
	for i := 0; i < totalNode; i++ {
		distance[i] = 2147483647
	}

	pq := &MinDistanceHeap{}
	heap.Init(pq)

	distance[source] = 0
	heap.Push(pq, GraphNode{distance: 0, node: source})

	for pq.Len() != 0 {

		pop := pq.Pop()

		nodeDistance := pop.(GraphNode).distance
		popNode := pop.(GraphNode).node

		for _, neighbour := range adjList[popNode] {

			if nodeDistance+neighbour.weight < distance[neighbour.v] {

				// find record already present
				var recordIndex int = -1
				for index, gn := range *pq {
					if gn.distance == distance[neighbour.v] && gn.node == neighbour.v {
						recordIndex = index
						break
					}
				}

				// if record found
				if recordIndex != -1 {
					heap.Remove(pq, recordIndex)
				}

				// distance update
				distance[neighbour.v] = nodeDistance + neighbour.weight

				// record push in pq
				heap.Push(pq, GraphNode{distance: distance[neighbour.v], node: neighbour.v})
			}
		}
	}

	//ans := []int{}
	//ans = append(ans, distance...)

	fmt.Println(distance)

}

type GraphNode struct {
	distance int
	node     int
}

type MinDistanceHeap []GraphNode

func (h MinDistanceHeap) Len() int {
	return len(h)
}

func (h MinDistanceHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinDistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(GraphNode))
}

func (h *MinDistanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
