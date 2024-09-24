package linkedlist

func mergeTwoSortedLists(first *Node, second *Node) *Node {

	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	var mergeNode *Node
	if first.Data <= second.Data {
		mergeNode = mergeSortedLinkedList(first, second)
	} else {
		mergeNode = mergeSortedLinkedList(second, first)
	}

	return mergeNode
}

func mergeSortedLinkedList(first, second *Node) *Node {

	// If first list have only one node, point its next to entire second list
	if first.Next == nil {
		first.Next = second
		return first
	}

	curr1 := first
	next1 := first.Next

	curr2 := second
	next2 := second.Next

	for next1 != nil && curr2 != nil {

		if (curr2.Data >= curr1.Data) && (curr2.Data <= next1.Data) {
			// add node in between the node of first list
			curr1.Next = curr2
			next2 = curr2.Next
			curr2.Next = next1

			// updating pointer
			curr1 = curr2
			curr2 = next2
		} else {
			curr1 = next1
			next1 = next1.Next

			if next1 == nil {
				curr1.Next = curr2
				return first
			}
		}
	}

	return first

}
