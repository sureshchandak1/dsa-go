package linkedlist

import "fmt"

type CloneNode struct {
	data   int
	next   *CloneNode
	randum *CloneNode
}

type cloneLinkedList1 struct {
	head    *CloneNode
	tail    *CloneNode
	mapping map[*CloneNode]*CloneNode
}

func (c *cloneLinkedList1) clone(head *CloneNode) {

	c.mapping = make(map[*CloneNode]*CloneNode)

	// Clone normal node
	curr := head
	for curr != nil {

		// Create new node for clone
		c.insertAtTail(curr)

		curr = curr.next
	}

	// Clone randum node
	curr = head
	for curr != nil {

		newNode := c.mapping[curr]
		newNode.randum = c.mapping[curr.randum]

		curr = curr.next
	}

	c.printLinkedList()

}

func (c *cloneLinkedList1) insertAtTail(originalNode *CloneNode) {

	newNode := CloneNode{data: originalNode.data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
	} else {
		c.tail.next = &newNode
		c.tail = &newNode
	}

	c.mapping[originalNode] = &newNode
}

func (c *cloneLinkedList1) printLinkedList() {
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
