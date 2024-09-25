package linkedlist

import (
	"fmt"
	"testing"
)

func TestSorto12(t *testing.T) {

	var first = SingleLinkedList{}

	first.insert(1)
	first.insert(2)
	first.insert(0)
	first.insert(2)
	first.insert(1)
	first.insert(0)
	first.printLinkedList()

	sortNode := sort012(first.head)

	// print linked list
	fmt.Print("[ ")
	var currNode = sortNode
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
