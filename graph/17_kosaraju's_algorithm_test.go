package graph

import "testing"

func TestKosarajuAlgo(t *testing.T) {

	edges := [][]int{{0, 1}, {1, 2}, {1, 4}, {2, 3}, {3, 2}, {4, 0}}
	kosarajuAlgo(5, edges)

	edges = [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 2}, {3, 0}, {4, 3}, {3, 4}, {5, 0},
		{4, 5}, {7, 4}, {5, 6}, {6, 5}, {7, 6}}
	kosarajuAlgo(8, edges)

	edges = [][]int{{1, 0}, {0, 2}, {2, 1}, {3, 2}, {4, 3}, {2, 4}}
	kosarajuAlgo(5, edges)

}
