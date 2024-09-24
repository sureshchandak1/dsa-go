package linkedlist

func findMiddleNode(head *Node) *Node {

	length := getLength(head)

	mid := length/2 + 1

	currNode := head
	count := 1
	for count < mid {
		currNode = currNode.Next
		count++
	}

	return currNode
}

// two-pointer approach
// Optimized solution
func findMiddleNodeOptimized(head *Node) *Node {

	if head == nil || head.Next == nil {
		return head
	}

	// only 2 node in linked list
	if head.Next.Next == nil {
		return head.Next
	}

	slow := head
	fast := head

	for fast != nil {
		// fast move 2 step
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		// slow move 1 step
		slow = slow.Next
	}

	return slow

}

func getLength(head *Node) int {

	currNode := head
	length := 0

	for currNode != nil {
		currNode = currNode.Next
		length++
	}

	return length
}
