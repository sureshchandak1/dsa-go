package linkedlist

func (s *SingleLinkedList) reverseLinkedList(head *Node) (*Node, int, string) {

	if head == nil || head.Next == nil {
		return head, 0, ""
	}

	var prev *Node = nil
	var curr *Node = head
	var forward *Node = nil

	for curr != nil {
		forward = curr.Next

		curr.Next = prev

		prev = curr
		curr = forward
	}

	s.head = prev
	return prev, 0, ""
}
