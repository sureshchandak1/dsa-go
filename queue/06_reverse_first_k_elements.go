package queue

import "container/list"

func reverseKElements(queue *list.List, k int) {

	// Step 1: pop first k from Q and put into stack
	stack := list.New()

	for i := 0; i < k; i++ {
		front := queue.Front()
		stack.PushBack(front.Value)
		queue.Remove(front)
	}

	// Step 2: fetch from stack and push into q
	for stack.Len() != 0 {
		back := stack.Back()
		queue.PushBack(back.Value)
		stack.Remove(back)
	}

	// Step 3: fetch from (n - k) element from Q and push back
	t := queue.Len() - k - 1

	for t >= 0 {
		front := queue.Front()
		queue.Remove(front)
		queue.PushBack(front.Value)

		t--
	}
}
