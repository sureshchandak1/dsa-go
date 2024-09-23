package linkedlist

import "fmt"

type SingleLinkedList struct {
	head *Node
	tail *Node
}

func (s *SingleLinkedList) insert(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {

		var lastNode = s.head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		lastNode.Next = &newNode
		s.tail = &newNode
	}
}

func (s *SingleLinkedList) insertAtHead(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {
		newNode.Next = s.head
		s.head = &newNode
	}
}

func (s *SingleLinkedList) insertAtTail(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {
		s.tail.Next = &newNode
		s.tail = &newNode
	}
}

func (s *SingleLinkedList) insertAtPosition(data int, position int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {

		if position == 1 {
			s.insertAtHead(data)
			return
		}

		currNode := s.head
		count := 1
		for count < position-1 && currNode.Next != nil {
			currNode = currNode.Next
			count++
		}

		if currNode.Next == nil {
			s.insertAtTail(data)
			return
		}

		nextNode := currNode.Next
		currNode.Next = &newNode
		newNode.Next = nextNode
	}
}

func (s *SingleLinkedList) deleteAtPosition(position int) {

	if s.head == nil {
		return
	}

	if position == 1 {
		s.head = s.head.Next
	}

	previousNode := s.head
	currentNode := s.head
	count := 1
	for count < position {
		previousNode = currentNode
		currentNode = currentNode.Next
		count++
	}

	if currentNode.Next == nil {
		previousNode.Next = nil
		s.tail = previousNode
	} else {
		previousNode.Next = currentNode.Next
	}
}

func (s *SingleLinkedList) size() int {

	currNode := s.head
	length := 0

	for currNode != nil {
		currNode = currNode.Next
		length++
	}

	return length
}

func (s *SingleLinkedList) printLinkedList() {
	fmt.Print("[ ")
	var currNode = s.head
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
