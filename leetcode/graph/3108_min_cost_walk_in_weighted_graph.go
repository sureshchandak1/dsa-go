package graph

import "math"

// disjoint set (union find algorithm)
type DSU struct {
	rank   []int
	parent []int
}

func (s *DSU) init(n int) {
	s.rank = make([]int, n)
	s.parent = make([]int, n)

	for i := range n {
		s.parent[i] = i
		s.rank[i] = 0
	}
}

func (s *DSU) findParent(node int) int {

	// Base case
	if node == s.parent[node] {
		return node
	}

	s.parent[node] = s.findParent(s.parent[node])

	return s.parent[node]

}

func (s *DSU) union(node1, node2 int) {

	rootParent1 := s.findParent(node1)
	rootParent2 := s.findParent(node2)

	if rootParent1 == rootParent2 {
		return
	}

	if s.rank[rootParent1] < s.rank[rootParent2] {
		s.parent[rootParent1] = rootParent2
	} else if s.rank[rootParent2] < s.rank[rootParent1] {
		s.parent[rootParent2] = rootParent1
	} else {
		s.parent[rootParent2] = rootParent1
		s.rank[rootParent1]++
	}
}

func minimumCostWalkInWeightedGraph(n int, edges [][]int, query [][]int) []int {
	var d DSU

	d.init(n)

	for _, edge := range edges {
		d.union(edge[0], edge[1])
	}

	componentConst := make([]int, n)
	intMax := int(math.Pow(2, 31) - 1)
	for i := range n {
		componentConst[i] = intMax
	}

	for _, edge := range edges {
		rootParent := d.findParent(edge[0])
		componentConst[rootParent] = componentConst[rootParent] & edge[2]
	}

	m := len(query)
	result := make([]int, m)

	for i := range m {
		rootParent1 := d.findParent(query[i][0])
		rootParent2 := d.findParent(query[i][1])

		if rootParent1 != rootParent2 {
			result[i] = -1
		} else {
			result[i] = componentConst[rootParent1]
		}
	}

	return result
}
