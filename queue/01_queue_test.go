package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := newQueue(10)
	queue.push(10)
	queue.push(20)
	queue.push(30)
	queue.push(40)
	queue.push(50)

	fmt.Println(queue.arr)

	fmt.Println("First:", queue.first())

	fmt.Println("Pop:", queue.pop())

	fmt.Println("First:", queue.first())

	fmt.Println("Pop:", queue.pop())

	queue.push(60)
	queue.push(70)

	fmt.Println(queue.arr)
}
