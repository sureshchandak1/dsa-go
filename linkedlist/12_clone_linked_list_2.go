package linkedlist

import "fmt"

type cloneLinkedList2 struct {
	head *CloneNode
	tail *CloneNode
}

func (c *cloneLinkedList2) clone(head *CloneNode) {

	// Step 1: Create clone list
	curr := head

	for curr != nil {
		c.insertAtTail(curr)

		curr = curr.next
	}

	// Step 2: clone nodes add in between original list
	originalNode := head
	cloneNode := c.head
	var next *CloneNode = nil

	for originalNode != nil && cloneNode != nil {
		next = originalNode.next
		originalNode.next = cloneNode
		originalNode = next

		next = cloneNode.next
		cloneNode.next = originalNode
		cloneNode = next
	}

	// Step 3: random pointer copy
	originalNode = head

	for originalNode != nil {

		cloneNode := originalNode.next
		if cloneNode != nil {
			if originalNode.randum != nil {
				cloneNode.randum = originalNode.randum.next
			} else {
				cloneNode.randum = originalNode.randum
			}
		}

		originalNode = originalNode.next.next
	}

	// Step 4: revert changes done in step 2
	originalNode = head
	cloneNode = c.head

	for originalNode != nil && cloneNode != nil {
		originalNode.next = cloneNode.next
		originalNode = originalNode.next

		if originalNode != nil {
			cloneNode.next = originalNode.next
		}

		cloneNode = cloneNode.next
	}

	// Step 5: print ans
	c.printLinkedList()

}

func (c *cloneLinkedList2) insertAtTail(originalNode *CloneNode) {

	newNode := CloneNode{data: originalNode.data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
	} else {
		c.tail.next = &newNode
		c.tail = &newNode
	}

}

func (c *cloneLinkedList2) printLinkedList() {
	fmt.Print("[ ")
	var currNode = c.head
	for currNode != nil {
		fmt.Print(currNode.data)
		fmt.Print(" ")

		currNode = currNode.next
	}
	fmt.Print("]")
	fmt.Println()
}
