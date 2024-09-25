package linkedlist

func (s *SingleLinkedList) isPalindrome() bool {

	if s.head == nil || s.head.Next == nil {
		return true
	}

	// Step - 1 : Find middle
	var middleNode *Node = getMid(s.head)

	// Step - 2 : Reverse list after middle
	var temp *Node = middleNode.Next
	middleNode.Next = reverseLinkedList(temp)

	// Step - 3 : Compare both sides
	head1 := s.head
	head2 := middleNode.Next

	for head2 != nil {

		if head1.Data != head2.Data {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	// Step - 4 : Repeat step 2
	temp = middleNode.Next
	middleNode.Next = reverseLinkedList(temp)

	return true
}

func getMid(head *Node) *Node {

	slow := head
	fast := head.Next

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		slow = slow.Next
	}

	return slow
}

func reverseLinkedList(head *Node) *Node {

	var prev *Node = nil
	var curr *Node = head
	var next *Node = nil

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
