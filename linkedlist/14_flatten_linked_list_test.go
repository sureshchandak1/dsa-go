package linkedlist

import (
	"fmt"
	"testing"
)

func TestFlattenLinkedList(t *testing.T) {

	node1 := &flattenNode{data: 1}
	/*node2 := &flattenNode{data: 4}
	node3 := &flattenNode{data: 7}
	node4 := &flattenNode{data: 9}
	node5 := &flattenNode{data: 20}

	child1 := &flattenNode{data: 2}
	child1.next = &flattenNode{data: 3}
	node1.child = child1

	child2 := &flattenNode{data: 5}
	child2.next = &flattenNode{data: 6}
	node2.child = child2

	node3.child = &flattenNode{data: 8}
	node4.child = &flattenNode{data: 12}

	node5.child = nil
	node5.next = nil*/

	result := flattenLinkedList(node1)

	// print sorted list
	fmt.Print("[ ")
	var currNode = result
	for currNode != nil {
		fmt.Print(currNode.data)
		fmt.Print(" ")

		currNode = currNode.next
	}
	fmt.Print("]")
	fmt.Println()
}
