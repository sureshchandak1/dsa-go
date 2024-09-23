package linkedlist

import (
	"fmt"
	"testing"
)

func TestSingleLinkedList(t *testing.T) {

	var linkedlist LinkedList = &SingleLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.printLinkedList()
	linkedlist.insertAtHead(9)
	linkedlist.printLinkedList()
	linkedlist.insertAtTail(40)
	linkedlist.printLinkedList()
	linkedlist.insertAtPosition(25, 4)
	linkedlist.insertAtPosition(50, 10)
	linkedlist.printLinkedList()
	linkedlist.deleteAtPosition(4)

	linkedlist.printLinkedList()
	fmt.Println("Size:", linkedlist.size())
}

func TestDoublyLinkedList(t *testing.T) {

	var linkedlist LinkedList = &DoublyLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.printLinkedList()
	linkedlist.insertAtHead(9)
	linkedlist.printLinkedList()
	linkedlist.insertAtTail(40)
	linkedlist.printLinkedList()
	linkedlist.insertAtPosition(25, 4)
	linkedlist.insertAtPosition(50, 10)
	linkedlist.printLinkedList()
	linkedlist.deleteAtPosition(4)

	linkedlist.printLinkedList()
	fmt.Println("Size:", linkedlist.size())
}

func TestCircularLinkedList(t *testing.T) {

	var linkedlist LinkedList = &CircularLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.printLinkedList()
	linkedlist.insertAtHead(9)
	linkedlist.printLinkedList()
	linkedlist.insertAtTail(40)
	linkedlist.printLinkedList()
	linkedlist.insertAtPosition(25, 4)
	linkedlist.insertAtPosition(50, 10)
	linkedlist.printLinkedList()
	linkedlist.deleteAtPosition(4)

	linkedlist.printLinkedList()
	fmt.Println("Size:", linkedlist.size())
}
