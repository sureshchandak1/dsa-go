package linkedlist

func removeLoop(head *Node) {

	if head == nil {
		return
	}

	startOfLoop := getLoopStartingNode(head)

	temp := startOfLoop

	for temp.Next != startOfLoop {
		temp = temp.Next
	}

	temp.Next = nil
}

func getLoopStartingNode(head *Node) *Node {

	if head == nil {
		return nil
	}

	intersection := floydDetectLoop(head)
	slow := head

	for slow != intersection {
		slow = slow.Next
		intersection = intersection.Next
	}

	return slow
}

func floydDetectLoop(head *Node) *Node {

	if head == nil {
		return nil
	}

	fast := head // move 2 step
	slow := head // move 1 step

	for fast != nil && slow != nil {

		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		slow = slow.Next

		// detect loop is present
		if fast == slow {
			return slow
		}
	}

	return nil
}

func detectLoop(head *Node) bool {

	if head == nil {
		return false
	}

	visited := make(map[*Node]bool)

	curr := head

	for curr != nil {

		// If already visited then loop is present
		if visited[curr] {
			return true
		}

		visited[curr] = true

		curr = curr.Next
	}

	return false
}
