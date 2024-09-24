package linkedlist

// Remove duplicates from a sorted/unsorted Linked List

func removeDuplicatesSorted(head *Node) {

	if head == nil {
		return
	}

	curr := head

	for curr != nil {

		next := curr.Next
		for next != nil && curr.Data == next.Data {
			next = next.Next
		}

		curr.Next = next
		curr = next
	}

}

func removeDuplicatesUnsorted(head *Node) {

	if head == nil {
		return
	}

	visited := make(map[int]bool)

	var prev *Node = nil
	var curr *Node = head

	for curr != nil {

		if visited[curr.Data] {
			prev.Next = curr.Next
		} else {
			visited[curr.Data] = true
			prev = curr
		}

		curr = curr.Next
	}
}
