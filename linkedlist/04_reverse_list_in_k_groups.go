package linkedlist

func (s *SingleLinkedList) reverseKGroup(head *Node, k int) *Node {

	if head == nil || head.Next == nil {
		return head
	}

	length := getLength(head)
	if k > length {
		return head
	}

	var prev *Node = nil
	var curr *Node = head
	var next *Node = nil

	count := 0

	for curr != nil && count < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next

		count++
	}

	if next != nil {
		// Recursive call for next k group
		head.Next = s.reverseKGroup(next, k)
	}

	s.head = prev
	return prev
}
