package graph

import "testing"

func TestGraphCycleDirectedGraph(t *testing.T) {

	edges := [][]int{{1, 2}, {4, 1}, {2, 4}, {3, 4}, {5, 2}, {1, 3}}
	isCyclicDFSDirectedGraph(5, edges)

	edges = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	isCyclicDFSDirectedGraph(5, edges)

}
