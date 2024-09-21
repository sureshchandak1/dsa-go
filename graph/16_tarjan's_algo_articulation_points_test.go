package graph

import "testing"

func TestArticulationPointsAlgo(t *testing.T) {

	edges := [][]int{{0, 3}, {3, 4}, {0, 4}, {0, 1}, {1, 2}}
	articulationPointsAlgo(5, edges)

}
