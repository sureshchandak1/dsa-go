package linkedlist

import "testing"

func TestAddTwoLinkedList(t *testing.T) {

	var first = SingleLinkedList{}

	first.insert(1)
	first.insert(0)
	first.insert(0)
	first.printLinkedList()

	var second = SingleLinkedList{}

	second.insert(2)
	second.insert(0)
	second.insert(0)
	second.printLinkedList()

	addTwoLinkedList(first.head, second.head)

}
