package linkedlist

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type LinkedList interface {
	insert(data int)
	insertAtHead(data int)
	insertAtTail(data int)
	insertAtPosition(data int, position int)
	deleteAtPosition(position int)
	size() int
	printLinkedList()
}
