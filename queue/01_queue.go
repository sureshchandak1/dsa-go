package queue

type queue struct {
	arr   []int
	front int
	rear  int
	size  int
}

func newQueue(size int) queue {
	return queue{
		arr:   make([]int, size),
		front: 0,
		rear:  0,
		size:  size,
	}
}

func (q *queue) push(data int) {
	if q.rear == q.size {
		return
	}

	q.arr[q.rear] = data
	q.rear++
}

func (q *queue) pop() int {
	if q.front == q.rear {
		return -1
	}

	value := q.arr[q.front]
	q.arr[q.front] = -1
	q.front++

	if q.front == q.rear {
		q.front = 0
		q.rear = 0
	}

	return value
}

func (q *queue) first() int {
	if q.front == q.rear {
		return -1
	}

	return q.arr[q.front]
}
