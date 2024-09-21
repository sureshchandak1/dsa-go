package graph

import "fmt"

func disjointUnionSets(totalNode int, edges [][]int) {

	parent := make([]int, totalNode)
	rank := make([]int, totalNode)

	for i := 0; i < totalNode; i++ {
		parent[i] = i // Initially, all elements are in their own set(parent)
		rank[i] = 0
	}

	for _, edge := range edges {

		u := edge[0]
		v := edge[1]

		unionSet(u, v, parent, rank)
	}

	fmt.Println(findParent(4, parent, rank) == findParent(0, parent, rank))
	fmt.Println(findParent(1, parent, rank) == findParent(0, parent, rank))
}

func unionSet(u, v int, parent, rank []int) {

	uParent := findParent(u, parent, rank)
	vParent := findParent(v, parent, rank)

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

func findParent(node int, parent, rank []int) int {

	// if x is not the parent of itself
	if parent[node] != node {
		// recursively call Find on its parent
		parent[node] = findParent(parent[node], parent, rank)
	}

	return parent[node]
}
