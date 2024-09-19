package graph

import "testing"

func TestShortestPathUndirectedGraph(t *testing.T) {

	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	shortestPathUndirectedGraph(edges, 1, 4)

	edges = [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 5}, {5, 8}, {3, 8}, {4, 6}, {6, 7}, {7, 8}}
	shortestPathUndirectedGraph(edges, 1, 8)

}
