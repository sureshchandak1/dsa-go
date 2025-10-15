package practice

import (
	"container/list"
	"sort"
)

func getAncestorsOptimized(n int, edges [][]int) [][]int {

	adjList := make([][]int, n)
	result := make([][]int, n)
	set := make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		adjList[i] = []int{}
		result[i] = []int{}
		set[i] = make(map[int]struct{})
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
	}

	topoSort := topoSortByKahnsAlgo(n, adjList)

	for _, node := range topoSort {
		for _, neighbour := range adjList[node] {
			set[neighbour][node] = struct{}{} // direct parent
			for grandParent := range set[node] {
				set[neighbour][grandParent] = struct{}{} // parent of parent
			}
		}
	}

	for i := 0; i < n; i++ {
		// Extract keys and sort them to maintain order similar to TreeSet
		keys := make([]int, 0, len(set[i]))
		for val := range set[i] {
			keys = append(keys, val)
		}
		sort.Ints(keys)
		result[i] = append(result[i], keys...)
	}

	return result
}

func topoSortByKahnsAlgo(n int, adjList [][]int) []int {

	// find indegree
	inDegree := make([]int, n)
	for _, neighbours := range adjList {
		for _, neighbour := range neighbours {
			inDegree[neighbour]++
		}
	}

	// BFS

	queue := list.New()
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	topoSort := make([]int, n)
	index := 0
	for queue.Len() > 0 {
		val := queue.Remove(queue.Front()).(int)
		topoSort[index] = val
		index++

		for _, node := range adjList[val] {
			inDegree[node]--
			if inDegree[node] == 0 {
				queue.PushBack(node)
			}
		}
	}

	return topoSort
}

func getAncestors(n int, edges [][]int) [][]int {

	adjList := make([][]int, n)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		adjList[i] = []int{}
		result[i] = []int{}
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjList[u] = append(adjList[u], v)
	}

	for node := 0; node < n; node++ {
		visited := make(map[int]struct{})
		visited[node] = struct{}{}
		dfsForAncestors(node, visited, result, node, adjList)
	}

	return result
}

func dfsForAncestors(currentNode int, visited map[int]struct{}, result [][]int, parent int, adjList [][]int) {

	for _, neighbour := range adjList[currentNode] {
		if _, ok := visited[neighbour]; !ok {
			visited[neighbour] = struct{}{}
			result[neighbour] = append(result[neighbour], parent)
			dfsForAncestors(neighbour, visited, result, parent, adjList)
		}
	}
}
