package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head *Node // head of list (Starting node)
	tail *Node // end of list (Ending node)
}

func (d *DoublyLinkedList) insert(data int) {

	newNode := Node{Data: data}

	if d.head == nil {
		d.head = &newNode
		d.tail = &newNode
	} else {

		lastNode := d.head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		newNode.Prev = lastNode
		lastNode.Next = &newNode
		d.tail = &newNode
	}
}

func (d *DoublyLinkedList) insertAtHead(data int) {

	newNode := Node{Data: data}

	if d.head == nil {
		d.head = &newNode
		d.tail = &newNode
	} else {
		d.head.Prev = &newNode
		newNode.Next = d.head
		d.head = &newNode
	}
}

func (d *DoublyLinkedList) insertAtTail(data int) {

	newNode := Node{Data: data}

	if d.head == nil {
		d.head = &newNode
		d.tail = &newNode
	} else {
		d.tail.Next = &newNode
		newNode.Prev = d.tail
		d.tail = &newNode
	}
}

func (d *DoublyLinkedList) insertAtPosition(data int, position int) {

	newNode := Node{Data: data}

	if d.head == nil {
		d.head = &newNode
		d.tail = &newNode
	} else {

		if position == 1 {
			d.insertAtHead(data)
			return
		}

		currNode := d.head
		count := 1
		for count < position-1 && currNode.Next != nil {
			currNode = currNode.Next
			count++
		}

		if currNode.Next == nil {
			d.insertAtTail(data)
			return
		}

		nextNode := currNode.Next
		currNode.Next = &newNode
		newNode.Prev = currNode
		newNode.Next = nextNode
		nextNode.Prev = &newNode
	}
}

func (d *DoublyLinkedList) deleteAtPosition(position int) {

	if d.head == nil {
		return
	}
	if position == 1 {
		d.head = d.head.Next
		d.head.Prev = nil
	}

	prevNode := d.head
	currNode := d.head
	count := 1
	for count < position {
		prevNode = currNode
		currNode = currNode.Next
		count++
	}

	if currNode.Next == nil {
		prevNode.Next = nil
		d.tail = prevNode
	} else {
		prevNode.Next = currNode.Next
		currNode.Next.Prev = prevNode
	}

}

func (d *DoublyLinkedList) size() int {

	if d.head == nil {
		return 0
	}

	currNode := d.head
	length := 1

	for currNode.Next != nil {
		currNode = currNode.Next
		length++
	}

	return length
}

func (d *DoublyLinkedList) printLinkedList() {

	fmt.Print("[ ")

	currNode := d.head
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
