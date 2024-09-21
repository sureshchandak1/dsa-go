package graph

import (
	"fmt"
	"sort"
)

func kruskalAlgoMST(totalNode int, edges [][]int) {

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	parent := make([]int, totalNode)
	rank := make([]int, totalNode)

	// Insert default value
	for i := 0; i < totalNode; i++ {
		parent[i] = i
		rank[i] = 0
	}

	minWeight := 0

	for i := 0; i < totalNode; i++ {

		uParent := findParentMST(edges[i][0], parent, rank)
		vParent := findParentMST(edges[i][1], parent, rank)
		wt := edges[i][2]

		if uParent != vParent {
			minWeight += wt
			unionSetMST(uParent, vParent, parent, rank)
		}
	}

	fmt.Println("Min weight:", minWeight)
}

func unionSetMST(u, v int, parent, rank []int) {

	uParent := findParentMST(u, parent, rank)
	vParent := findParentMST(v, parent, rank)

	// elements are in the same set, no need to unite anything
	if uParent == vParent {
		return
	}

	// find rank
	uRank := rank[uParent]
	vRank := rank[vParent]

	if uRank < vRank {
		parent[uParent] = vParent
	} else if vRank < uRank {
		parent[vParent] = uParent
	} else {
		parent[vParent] = uParent // then move v under u (doesn't matter which one goes where)
		rank[uParent]++           // increment the result tree's rank by 1
	}
}

func findParentMST(node int, parent, rank []int) int {

	// if x is not the parent of itself
	if parent[node] != node {
		// recursively call Find on its parent
		parent[node] = findParent(parent[node], parent, rank)
	}

	return parent[node]
}
