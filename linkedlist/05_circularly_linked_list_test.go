package linkedlist

import (
	"fmt"
	"testing"
)

func TestIsCircular(t *testing.T) {

	testIsCircular1()
	testIsCircular2()
}

func testIsCircular1() {

	var linkedlist SingleLinkedList = SingleLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.insertAtHead(9)
	linkedlist.insertAtTail(40)
	linkedlist.insertAtPosition(25, 4)
	linkedlist.insertAtPosition(50, 10)
	linkedlist.deleteAtPosition(4)

	fmt.Println("is circular:", isCircular(linkedlist.head))

}

func testIsCircular2() {

	var linkedlist CircularLinkedList = CircularLinkedList{
		head: nil,
		tail: nil,
	}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(30)
	linkedlist.insertAtHead(9)
	linkedlist.insertAtTail(40)
	linkedlist.insertAtPosition(25, 4)
	linkedlist.insertAtPosition(50, 10)
	linkedlist.deleteAtPosition(4)

	fmt.Println("is circular:", isCircular(linkedlist.head))

}
