package graph

import "testing"

func TestKruskalAlgoMST(t *testing.T) {

	edges := [][]int{{0, 1, 3}, {0, 3, 5}, {1, 2, 1}, {2, 3, 8}}
	kruskalAlgoMST(4, edges)

	edges = [][]int{{0, 1, 3}, {0, 2, 1}, {0, 3, 2}, {1, 2, 5}, {1, 3, 2}, {2, 3, 4}}
	kruskalAlgoMST(4, edges)

}
