package linkedlist

import (
	"fmt"
	"testing"
)

func TestCloneLinkedList1(t *testing.T) {

	node1 := CloneNode{data: 1}
	node2 := CloneNode{data: 2}
	node3 := CloneNode{data: 3}
	node4 := CloneNode{data: 4}
	node5 := CloneNode{data: 5}
	fmt.Printf("Original head address: %p\n", &node1)

	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5

	node1.randum = &node3
	node2.randum = &node1
	node3.randum = &node5
	node4.randum = &node3
	node5.randum = &node2

	cll := cloneLinkedList1{}

	cll.clone(&node1)

	fmt.Printf("Clone head address: %p\n", &cll.head)
}

func TestCloneLinkedList2(t *testing.T) {

	node1 := CloneNode{data: 1}
	node2 := CloneNode{data: 2}
	node3 := CloneNode{data: 3}
	node4 := CloneNode{data: 4}
	node5 := CloneNode{data: 5}
	fmt.Printf("Original head address: %p\n", &node1)

	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5

	node1.randum = &node3
	node2.randum = &node1
	node3.randum = &node5
	node4.randum = &node3
	node5.randum = &node2

	cll := cloneLinkedList2{}

	cll.clone(&node1)

	fmt.Printf("Clone head address: %p\n", &cll.head)

}
