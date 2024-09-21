package graph

import "testing"

func TestShortestPathDAG(t *testing.T) {

	edges := [][]int{{0, 1, 5}, {0, 2, 3}, {1, 2, 2}, {1, 3, 6}, {2, 3, 7}, {2, 4, 4}, {2, 5, 2}, {3, 4, -1}, {4, 5, -2}}
	shortestPathDAG(6, edges, 1)

}