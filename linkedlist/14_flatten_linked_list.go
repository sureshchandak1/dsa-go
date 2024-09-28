package linkedlist

type flattenNode struct {
	data  int
	next  *flattenNode
	child *flattenNode
}

func flattenLinkedList(head *flattenNode) *flattenNode {

	// Base case
	if head == nil || head.next == nil {
		return head
	}

	left := head
	right := flattenLinkedList(left.next)
	left.next = nil

	result := merge2Nodes(left, right)

	return result
}

func merge2Nodes(node1 *flattenNode, node2 *flattenNode) *flattenNode {

	if node1 == nil {
		return node2
	}

	if node2 == nil {
		return node1
	}

	var result *flattenNode = nil

	if node1.data < node2.data {
		result = node1
		result.child = merge2Nodes(node1.child, node2)
	} else {
		result = node2
		result.child = merge2Nodes(node1, node2.child)
	}

	result.next = nil
	return result

}
