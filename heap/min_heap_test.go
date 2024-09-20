package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {

	h := &MinHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	fmt.Printf("minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
