package graph

import (
	"fmt"
	"testing"
)

func TestMinimumCostWalkInWeightedGraph(t *testing.T) {
	fmt.Println(minimumCostWalkInWeightedGraph(5,
		[][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, [][]int{{0, 3}, {3, 4}}))
	fmt.Println(minimumCostWalkInWeightedGraph(3,
		[][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, [][]int{{1, 2}}))
	fmt.Println(minimumCostWalkInWeightedGraph(5,
		[][]int{{1, 3, 5}, {3, 1, 5}, {3, 0, 4}}, [][]int{{1, 2}, {3, 0}, {3, 1}, {2, 3}, {3, 0}, {0, 1}, {2, 3}, {1, 2}}))
}
