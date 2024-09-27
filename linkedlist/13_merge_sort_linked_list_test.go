package linkedlist

import (
	"fmt"
	"testing"
)

func TestMergeSortLinkedList(t *testing.T) {

	var linkedlist = SingleLinkedList{}

	linkedlist.insert(50)
	linkedlist.insert(2)
	linkedlist.insert(35)
	linkedlist.insert(25)
	linkedlist.insert(1)
	linkedlist.printLinkedList()

	result := mergeSort(linkedlist.head)

	// print sorted list
	fmt.Print("[ ")
	var currNode = result
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
