package hashmap

type MapNode[K any, V any] struct {
	key   K
	value V
}

func newNode[K any, V any](key K, value V) *MapNode[K, V] {
	return &MapNode[K, V]{key: key, value: value}
}

func init() {
	newNode(1, 10)

	var ll SingleLinkedList
	ll.insert(10)
	ll.insertAtTail(20)
	ll.insertAtHead(20)
}

type Node struct {
	Data int
	Next *Node
}

type SingleLinkedList struct {
	head *Node // head of list (Starting node)
	tail *Node // end of list (Ending node)
}

func (s *SingleLinkedList) insert(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {

		var lastNode = s.head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}

		lastNode.Next = &newNode
		s.tail = &newNode
	}
}

func (s *SingleLinkedList) insertAtHead(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {
		newNode.Next = s.head
		s.head = &newNode
	}
}

func (s *SingleLinkedList) insertAtTail(data int) {

	newNode := Node{Data: data}

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {
		s.tail.Next = &newNode
		s.tail = &newNode
	}
}
