package graph

import "testing"

func TestDFS(t *testing.T) {

	edges := [][]int{{0, 4}, {2, 0}, {2, 4}, {3, 0}, {3, 2}, {4, 3}}
	dfsTraversal(5, edges)

	edges = [][]int{{9, 0}, {8, 1}, {7, 2}, {6, 3}, {5, 4}}
	dfsTraversal(10, edges)
}
