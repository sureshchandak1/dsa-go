package queue

import (
	"container/list"
	"testing"
)

func TestInterleaveQueue(t *testing.T) {

	queue := list.New()
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)
	queue.PushBack(40)
	queue.PushBack(50)
	queue.PushBack(60)

	printList(queue)

	interleaveQueue(queue)

	printList(queue)

	queue = list.New()
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)
	queue.PushBack(40)
	queue.PushBack(50)
	queue.PushBack(60)

	printList(queue)

	interleaveQueueWithStack(queue)

	printList(queue)
}
