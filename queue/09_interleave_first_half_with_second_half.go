package queue

import "container/list"

func interleaveQueue(queue *list.List) {

	size := queue.Len()
	mid := size / 2

	newQueue := list.New()

	for i := 0; i < mid; i++ {
		front := queue.Front()
		queue.Remove(front)
		newQueue.PushBack(front.Value)
	}

	for newQueue.Len() != 0 {

		qFront := queue.Front()
		queue.Remove(qFront)

		nFront := newQueue.Front()
		newQueue.Remove(nFront)

		queue.PushBack(nFront.Value)
		queue.PushBack(qFront.Value)
	}
}

func interleaveQueueWithStack(queue *list.List) {

	size := queue.Len()
	mid := size / 2

	stack := list.New()

	for i := 0; i < mid; i++ {
		front := queue.Front()
		queue.Remove(front)
		stack.PushBack(front.Value)
	}

	reverseStack(stack)

	for stack.Len() != 0 {
		qFront := queue.Front()
		queue.Remove(qFront)

		bottom := stack.Back()
		stack.Remove(bottom)

		queue.PushBack(bottom.Value)
		queue.PushBack(qFront.Value)
	}
}

func reverseStack(stack *list.List) {

	// Base case
	if stack.Len() == 0 {
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive call
	reverseStack(stack)

	insertAtBottom(stack, back.Value.(int))
}

func insertAtBottom(stack *list.List, data int) {

	// Base case
	if stack.Len() == 0 {
		stack.PushBack(data)
		return
	}

	back := stack.Back()
	stack.Remove(back)

	// Recursive
	insertAtBottom(stack, data)

	stack.PushBack(back.Value)
}
