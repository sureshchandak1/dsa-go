package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedListMiddleNode(t *testing.T) {

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

	linkedlist.printLinkedList()
	fmt.Println("Size:", linkedlist.size())

	middleNode := findMiddleNode(linkedlist.head)
	fmt.Println("Middle Node:", middleNode.Data)

	middleNode = findMiddleNodeOptimized(linkedlist.head)
	fmt.Println("Middle Node:", middleNode.Data)

}
