package graph

import "testing"

func TestBFS(t *testing.T) {

	edges := [][]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}}
	bfsTraversal(edges)

	edges = [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}, {1, 7}, {2, 5}, {3, 6}}
	bfsTraversal(edges)

}
