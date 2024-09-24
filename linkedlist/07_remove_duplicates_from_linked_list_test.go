package linkedlist

import (
	"testing"
)

func TestRemoveDuplicatesSorted(t *testing.T) {

	var linkedlist = SingleLinkedList{}

	linkedlist.insert(1)
	linkedlist.insert(2)
	linkedlist.insert(3)
	linkedlist.insert(3)
	linkedlist.insert(3)
	linkedlist.insert(4)
	linkedlist.insert(5)
	linkedlist.insert(5)
	linkedlist.printLinkedList()

	removeDuplicatesSorted(linkedlist.head)

	linkedlist.printLinkedList()
}

func TestRemoveDuplicatesUnSorted(t *testing.T) {

	var linkedlist = SingleLinkedList{}

	linkedlist.insert(10)
	linkedlist.insert(20)
	linkedlist.insert(50)
	linkedlist.insert(30)
	linkedlist.insert(10)
	linkedlist.insert(40)
	linkedlist.insert(20)
	linkedlist.insert(50)
	linkedlist.printLinkedList()

	removeDuplicatesUnsorted(linkedlist.head)

	linkedlist.printLinkedList()
}
