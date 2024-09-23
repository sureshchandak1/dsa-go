package linkedlist

import "fmt"

type CircularLinkedList struct {
	head *Node // head of list (Starting node)
	tail *Node // end of list (Ending node)
}

func (c *CircularLinkedList) insert(data int) {

	newNode := Node{Data: data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
		c.tail.Next = &newNode
	} else {

		lastNode := c.head
		for lastNode.Next != c.head {
			lastNode = lastNode.Next
		}

		lastNode.Next = &newNode
		c.tail = &newNode
		c.tail.Next = c.head
	}
}

func (c *CircularLinkedList) insertAtHead(data int) {

	newNode := Node{Data: data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
		c.tail.Next = &newNode
	} else {
		newNode.Next = c.head
		c.tail.Next = &newNode
		c.head = &newNode
	}
}

func (c *CircularLinkedList) insertAtTail(data int) {

	newNode := Node{Data: data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
		c.tail.Next = &newNode
	} else {
		c.tail.Next = &newNode
		c.tail = &newNode
		c.tail.Next = c.head
	}
}

func (c *CircularLinkedList) insertAtPosition(data int, position int) {

	newNode := Node{Data: data}

	if c.head == nil {
		c.head = &newNode
		c.tail = &newNode
		c.tail.Next = &newNode
	} else {

		if position == 1 {
			c.insertAtHead(data)
			return
		}

		currNode := c.head
		count := 1
		for count < position-1 && currNode.Next != c.head {
			currNode = currNode.Next
			count++
		}

		if currNode.Next == c.head {
			c.insertAtTail(data)
			return
		}

		nextNode := currNode.Next
		currNode.Next = &newNode
		newNode.Next = nextNode
	}
}

func (c *CircularLinkedList) deleteAtPosition(position int) {

	if c.head == nil {
		return
	}
	if position == 1 {
		c.head = nil
		c.tail = nil
	}

	prevNode := c.head
	currNode := c.head
	count := 1
	for count < position {
		prevNode = currNode
		currNode = currNode.Next
		count++
	}

	if currNode.Next == c.head {
		prevNode.Next = currNode.Next
		c.tail = prevNode
	} else {
		prevNode.Next = currNode.Next
	}

}

func (c *CircularLinkedList) size() int {

	if c.head == nil {
		return 0
	}

	currNode := c.head
	length := 0
	for {
		length++

		if currNode.Next == c.head {
			break
		}

		currNode = currNode.Next
	}

	return length
}

func (c *CircularLinkedList) printLinkedList() {

	fmt.Print("[ ")
	currNode := c.head
	for {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		if currNode.Next == c.head {
			break
		}

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
