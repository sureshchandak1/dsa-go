package graph

import "testing"

func TestTopologicalSortDFS(t *testing.T) {

	edges := [][]int{{0, 3}, {2, 0}, {1, 0}, {1, 2}, {4, 2}, {4, 0}, {4, 3}}
	topologicalSortDFS(5, edges)

	edges = [][]int{{0, 2}, {3, 2}, {3, 1}, {0, 1}, {3, 0}, {2, 1}}
	topologicalSortDFS(4, edges)

}
