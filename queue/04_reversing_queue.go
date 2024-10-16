package queue

import "container/list"

func reverseQueue(queue *list.List) {

	stack := list.New()

	for queue.Len() != 0 {
		front := queue.Front()
		stack.PushBack(front.Value)
		queue.Remove(front)
	}

	for stack.Len() != 0 {
		top := stack.Back()
		queue.PushBack(top.Value)
		stack.Remove(top)
	}
}
