package graph

import "testing"

func TestKahnsAlgoTopoSort(t *testing.T) {

	edges := [][]int{{0, 3}, {2, 0}, {1, 0}, {1, 2}, {4, 2}, {4, 0}, {4, 3}}
	kahnsAlgoTopoSort(5, edges)

	edges = [][]int{{0, 2}, {3, 2}, {3, 1}, {0, 1}, {3, 0}, {2, 1}}
	kahnsAlgoTopoSort(4, edges)

}
