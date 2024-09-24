package linkedlist

import (
	"fmt"
	"testing"
)

func TestRemoveLoop(t *testing.T) {

	head := Node{Data: 10}
	head.Next = &Node{Data: 20}
	head.Next.Next = &Node{Data: 30}

	loopNode := Node{Data: 40}
	head.Next.Next.Next = &loopNode
	loopNode.Next = &Node{Data: 50}
	loopNode.Next.Next = &Node{Data: 60}
	// creating loop
	loopNode.Next.Next.Next = head.Next.Next

	fmt.Println("is loop present:", detectLoop(&head))
	fmt.Println("loop start node:", getLoopStartingNode(&head).Data)
	removeLoop(&head)
	fmt.Println("is loop present:", detectLoop(&head))
}
