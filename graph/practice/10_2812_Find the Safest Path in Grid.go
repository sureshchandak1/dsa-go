package practice

import (
	"container/heap"
)

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	// Base case, If thief is present at source or destination, then return 0
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}

	// Step 1: find the manhattan distance - matrix
	dist := findManhattanDist(grid, n)

	// Step 2: find the minimum popped path value
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Max heap
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{priority: dist[0][0], row: 0, col: 0})
	visited[0][0] = true

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(*Item)

		if cell.priority == 0 {
			return 0
		}

		if cell.row == n-1 && cell.col == n-1 {
			return cell.priority
		}

		// insert the neighbours
		for _, d := range dir {
			nextRow := cell.row + d[0]
			nextCol := cell.col + d[1]

			if isValid(n, nextRow, nextCol) && !visited[nextRow][nextCol] {
				minVal := dist[nextRow][nextCol]
				if cell.priority < minVal {
					minVal = cell.priority
				}
				heap.Push(pq, &Item{priority: minVal, row: nextRow, col: nextCol})
				visited[nextRow][nextCol] = true
			}
		}
	}

	return 0
}

// Multi Source BFS
func findManhattanDist(grid [][]int, n int) [][]int {
	type cell struct {
		dist int
		row  int
		col  int
	}

	queue := []cell{}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	// insert all the thief cells
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, cell{0, i, j})
				dist[i][j] = 0
			}
		}
	}

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	// apply Multi Source BFS
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for _, d := range dir {
			nextRow := top.row + d[0]
			nextCol := top.col + d[1]

			if isValid(n, nextRow, nextCol) && dist[nextRow][nextCol] == -1 {
				newDist := top.dist + 1
				dist[nextRow][nextCol] = newDist
				queue = append(queue, cell{newDist, nextRow, nextCol})
			}
		}
	}

	return dist
}

func isValid(n, row, col int) bool {
	return row >= 0 && row < n && col >= 0 && col < n
}

// PriorityQueue and Item for max heap implementation
type Item struct {
	priority int
	row      int
	col      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want a max heap, so use greater than
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
