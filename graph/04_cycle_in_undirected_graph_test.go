package graph

import "testing"

func TestGraphCycle(t *testing.T) {

	edges := [][]int{{1, 2}, {2, 3}}
	isCyclicBFSUndirectedGraph(3, edges)
	isCyclicDFSUndirectedGraph(3, edges)

	edges = [][]int{{1, 2}, {2, 3}, {4, 3}, {5, 4}, {6, 5}}
	isCyclicBFSUndirectedGraph(6, edges)
	isCyclicDFSUndirectedGraph(3, edges)

	edges = [][]int{{1, 4}, {3, 1}, {5, 4}, {5, 1}}
	isCyclicBFSUndirectedGraph(5, edges)
	isCyclicDFSUndirectedGraph(3, edges)

}
