package graph

import "testing"

func TestBridges(t *testing.T) {

	edges := [][]int{{0, 1}, {3, 1}, {1, 2}, {3, 4}}
	bridges(5, edges)

	edges = [][]int{{1, 2}, {1, 0}, {0, 2}, {0, 4}, {5, 4}, {5, 3}, {3, 4}}
	bridges(6, edges)

}
