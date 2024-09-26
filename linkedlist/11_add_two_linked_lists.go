package linkedlist

import "fmt"

type addLinkedList struct {
	head *Node
	tail *Node
}

func addTwoLinkedList(first *Node, second *Node) {

	llSum := addLinkedList{}

	// Step - 1 : reverse input Linked List
	first = llSum.reverse(first)
	second = llSum.reverse(second)

	// Step 2 : add 2 Linkedlist
	llSum.add(first, second)

	// Step 3
	sumNode := llSum.reverse(llSum.head)

	fmt.Print("Sum: ")
	llSum.printLinkedList(sumNode)

}

func (ll *addLinkedList) add(first *Node, second *Node) {

	carry := 0

	for first != nil || second != nil || carry != 0 {

		var value1 int = 0
		if first != nil {
			value1 = first.Data
		}

		var value2 int = 0
		if second != nil {
			value2 = second.Data
		}

		sum := value1 + value2 + carry

		digit := sum % 10

		ll.insertAtTail(digit)

		carry = sum / 10

		if first != nil {
			first = first.Next
		}

		if second != nil {
			second = second.Next
		}
	}

}

func (ll *addLinkedList) insertAtTail(value int) {

	newNode := Node{Data: value}

	if ll.head == nil {
		ll.head = &newNode
		ll.tail = &newNode
	} else {
		ll.tail.Next = &newNode
		ll.tail = &newNode
	}
}

func (ll *addLinkedList) reverse(head *Node) *Node {

	var prev *Node = nil
	var curr *Node = head
	var next *Node = nil

	for curr != nil {
		next = curr.Next

		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev

}

func (ll *addLinkedList) printLinkedList(head *Node) {
	fmt.Print("[ ")
	var currNode = head
	for currNode != nil {
		fmt.Print(currNode.Data)
		fmt.Print(" ")

		currNode = currNode.Next
	}
	fmt.Print("]")
	fmt.Println()
}
