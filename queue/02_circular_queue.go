package queue

type circularQueue struct {
	arr   []int
	front int
	rear  int
	size  int
}

func newCircularQueue(size int) circularQueue {
	return circularQueue{
		arr:   make([]int, size),
		front: -1,
		rear:  -1,
		size:  size,
	}
}

func (q *circularQueue) push(data int) bool {

	if (q.front == 0 && q.rear == q.size-1) || q.rear == (q.front-1)%(q.size-1) {
		return false // Queue is full
	}

	if q.front == -1 {
		// first element to push
		q.rear = 0
		q.front = 0
	} else if q.rear == q.size-1 && q.front != 0 {
		q.rear = 0 // to maintain cyclic nature
	} else {
		q.rear++
	}

	q.arr[q.rear] = data

	return true
}

func (q *circularQueue) pop() int {

	if q.front == -1 {
		return -1
	}

	value := q.arr[q.front]
	q.arr[q.front] = -1

	switch q.front {
	case q.rear:
		{
			// Single element in queue
			q.front = -1
			q.rear = -1
		}
	case q.size - 1:
		{
			q.front = 0 // to maintain cyclic nature
		}
	default:
		{
			q.front++
		}
	}

	return value
}

func (q *circularQueue) first() int {
	if q.front == q.rear {
		return -1
	}

	return q.arr[q.front]
}
