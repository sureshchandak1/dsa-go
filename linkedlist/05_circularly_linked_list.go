package linkedlist

func isCircular(head *Node) bool {

	if head == nil || head.Next == nil {
		return false
	}

	curr := head.Next

	for curr != nil && curr != head {
		curr = curr.Next
	}

	return curr == head

}
