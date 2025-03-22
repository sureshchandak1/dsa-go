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

func TestCountCompleteComponents(t *testing.T) {
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}))
	fmt.Println(countCompleteComponents(5, [][]int{{1, 2}, {3, 4}, {1, 4}, {2, 3}, {1, 3}, {2, 4}}))
}
