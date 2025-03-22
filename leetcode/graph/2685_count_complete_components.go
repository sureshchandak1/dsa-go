package graph

func countCompleteComponents(n int, edges [][]int) int {

	var dsu DSUWithSize

	dsu.init(n)

	edgeCount := make(map[int]int)

	for _, edge := range edges {
		dsu.union(edge[0], edge[1])
	}

	for _, edge := range edges {
		root := dsu.findParent(edge[0])
		edgeCount[root]++
	}

	completeCount := 0
	for vertex := range n {
		if dsu.findParent(vertex) == vertex {
			nodeCount := dsu.size[vertex]

			expectedEdges := (nodeCount * (nodeCount - 1)) / 2

			if edgeCount[vertex] == expectedEdges {
				completeCount++
			}
		}
	}

	return completeCount
}

// disjoint set (union find algorithm)
type DSUWithSize struct {
	parent []int
	size   []int
}

func (d *DSUWithSize) init(n int) {
	d.parent = make([]int, n)
	d.size = make([]int, n)

	for i := range n {
		d.parent[i] = i
		d.size[i] = 1
	}
}

func (d *DSUWithSize) findParent(node int) int {

	if d.parent[node] == node {
		return node
	}

	d.parent[node] = d.findParent(d.parent[node])
	return d.parent[node]
}

func (d *DSUWithSize) union(node1, node2 int) {

	rootParent1 := d.findParent(node1)
	rootParent2 := d.findParent(node2)

	if rootParent1 == rootParent2 {
		return
	}

	if d.size[rootParent1] > d.size[rootParent2] {
		d.parent[rootParent2] = rootParent1
		d.size[rootParent1] += d.size[rootParent2]
	} else {
		d.parent[rootParent1] = rootParent2
		d.size[rootParent2] += d.size[rootParent1]
	}
}
