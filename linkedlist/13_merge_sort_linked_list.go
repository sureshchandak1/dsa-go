package linkedlist

func mergeSort(head *Node) *Node {

	// Base case
	if head == nil || head.Next == nil {
		return head
	}

	// find linked list mid
	mid := linkedListMid(head)

	// break linked list into 2 parts
	left := head
	right := mid.Next
	mid.Next = nil

	// recursive calls to sort both parts
	left = mergeSort(left)
	right = mergeSort(right)

	// merge both left and right parts
	result := mergeNodes(left, right)

	return result
}

func linkedListMid(head *Node) *Node {

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func mergeNodes(left, right *Node) *Node {

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var ans *Node = &Node{Data: -1}
	var temp *Node = ans

	// merge 2 sorted linked list
	for left != nil && right != nil {
		if left.Data < right.Data {
			temp.Next = left
			temp = left
			left = left.Next
		} else {
			temp.Next = right
			temp = right
			right = right.Next
		}
	}

	for left != nil {
		temp.Next = left
		temp = left
		left = left.Next
	}

	for right != nil {
		temp.Next = right
		temp = right
		right = right.Next
	}

	return ans.Next
}
