package queue

type kQueue struct {
	arr      []int
	front    []int
	rear     []int
	next     []int
	freeSpot int
	size     int
}

func newKQueue(size int, totalQueue int) kQueue {

	kQueue := kQueue{
		arr:      make([]int, size),
		front:    make([]int, totalQueue),
		rear:     make([]int, totalQueue),
		next:     make([]int, size),
		freeSpot: 0,
		size:     size,
	}

	for i := 0; i < totalQueue; i++ {
		kQueue.front[i] = -1
		kQueue.rear[i] = -1
	}

	for i := 0; i < size; i++ {
		kQueue.next[i] = i + 1
	}
	kQueue.next[size-1] = -1

	return kQueue
}

func (q *kQueue) push(data int, qn int) {

	if q.freeSpot == -1 {
		return
	}

	index := q.freeSpot

	q.freeSpot = q.next[index]

	if q.front[qn-1] == -1 {
		// first time push
		q.front[qn-1] = index
	} else {
		// link new element to the previous element
		q.next[q.rear[qn-1]] = index
	}

	q.next[index] = -1

	q.rear[qn-1] = index

	q.arr[index] = data

}

func (q *kQueue) pop(qn int) int {

	if q.front[qn-1] == -1 {
		return -1
	}

	index := q.front[qn-1]

	q.front[qn-1] = q.next[index]

	q.next[index] = q.freeSpot

	q.freeSpot = index

	val := q.arr[index]
	q.arr[index] = -1

	return val
}
