package queue

type deque struct {
	arr   []int
	front int
	rear  int
	size  int
}

func newDeque(size int) deque {
	return deque{
		arr:   make([]int, size),
		front: -1,
		rear:  -1,
		size:  size,
	}
}

func (q *deque) pushFront(data int) bool {

	if q.isFull() {
		return false
	}

	if q.isEmpty() {
		// first element to push
		q.front = 0
		q.rear = 0
	} else if q.front == 0 && q.rear != q.size-1 {
		q.front = q.size - 1
	} else {
		q.front--
	}

	q.arr[q.front] = data

	return true
}

func (q *deque) pushRear(data int) bool {

	if q.isFull() {
		return false
	}

	if q.isEmpty() {
		// first element to push
		q.front = 0
		q.rear = 0
	} else if q.rear == q.size-1 && q.front != 0 {
		q.rear = 0 // to maintain cyclic nature
	} else {
		q.rear++
	}

	q.arr[q.rear] = data

	return true
}

func (q *deque) popFront() int {

	if q.isEmpty() {
		return -1
	}

	value := q.arr[q.front]
	q.arr[q.front] = -1

	if q.front == q.rear {
		// Single element in queue
		q.front = -1
		q.rear = -1
	} else if q.front == q.size-1 {
		q.front = 0 // to maintain cyclic nature
	} else {
		q.front++
	}

	return value
}

func (q *deque) popRear() int {

	if q.isEmpty() {
		return -1
	}

	value := q.arr[q.rear]
	q.arr[q.rear] = -1

	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	} else if q.rear == 0 {
		q.rear = q.size - 1 // to maintain cyclic nature
	} else {
		q.rear--
	}

	return value
}

func (q *deque) getFront() int {
	if q.isEmpty() {
		return -1
	}

	return q.arr[q.front]
}

func (q *deque) getRear() int {
	if q.isEmpty() {
		return -1
	}

	return q.arr[q.rear]
}

func (q *deque) isEmpty() bool {
	return q.front == -1
}

func (q *deque) isFull() bool {
	return (q.front == 0 && q.rear == q.size-1) ||
		(q.front != 0 && q.rear == (q.front-1)%(q.size-1))
}
