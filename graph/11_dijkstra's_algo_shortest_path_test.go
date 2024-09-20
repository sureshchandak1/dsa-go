package graph

import "testing"

func TestDijkstraShortPathAlgo(t *testing.T) {

	edges := [][]int{{0, 1, 5}, {0, 2, 8}, {1, 2, 9}, {1, 3, 2}, {2, 3, 6}}
	dijkstraShortPathAlgo(4, edges, 0)

	edges = [][]int{{0, 1, 7}, {0, 2, 1}, {0, 3, 2}, {1, 2, 3}, {1, 3, 5}, {1, 4, 1}, {3, 4, 7}}
	dijkstraShortPathAlgo(5, edges, 0)

}
