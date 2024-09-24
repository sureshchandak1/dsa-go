package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

	var first = SingleLinkedList{}

	first.insert(10)
	first.insert(20)
	first.insert(30)
	first.insert(40)
	first.printLinkedList()

	var second = SingleLinkedList{}

	second.insert(5)
	second.insert(15)
	second.insert(25)
	second.insert(35)
	second.printLinkedList()

	mergeNode := mergeTwoSortedLists(first.head, second.head)

	// print linked list
	fmt.Print("[ ")
	var currNode = mergeNode
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
