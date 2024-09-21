package graph

import "testing"

func TestDisjointUnionSets(t *testing.T) {

	edges := [][]int{{0, 2}, {4, 2}, {3, 1}}
	disjointUnionSets(5, edges)

}
