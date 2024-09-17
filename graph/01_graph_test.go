package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {

	addEdge(0, 1, 0)
	addEdge(0, 2, 0)
	addEdge(1, 2, 0)
	fmt.Println(adj)

	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 5}, {4, 6}}

	printAdjacency(edges)

}
