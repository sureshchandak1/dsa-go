package practice

import (
	"container/heap"
	"strings"
)

type Pair struct {
	count int
	ch    byte
}

type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count // max heap
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func longestDiverseString(a int, b int, c int) string {

	pq := &PriorityQueue{}

	if a > 0 {
		heap.Push(pq, &Pair{count: a, ch: 'a'})
	}
	if b > 0 {
		heap.Push(pq, &Pair{count: b, ch: 'b'})
	}
	if c > 0 {
		heap.Push(pq, &Pair{count: c, ch: 'c'})
	}

	var res strings.Builder
	for pq.Len() > 0 {

		node := heap.Pop(pq).(*Pair)
		n := res.Len()

		if n >= 2 && res.String()[n-1] == node.ch && res.String()[n-2] == node.ch {

			// If second highest element is not present then break
			if pq.Len() == 0 {
				break
			}

			sec := heap.Pop(pq).(*Pair)
			res.WriteByte(sec.ch)
			sec.count--
			if sec.count > 0 {
				heap.Push(pq, sec)
			}

		} else {
			res.WriteByte(node.ch)
			node.count--
		}

		// If element count is not 0, update count in pq
		if node.count > 0 {
			heap.Push(pq, node)
		}
	}

	return res.String()
}
