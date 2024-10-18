package queue

import (
	"fmt"
	"testing"
)

func TestKQueue(t *testing.T) {

	kQueue := newKQueue(10, 3)
	kQueue.push(10, 1)
	kQueue.push(15, 1)
	kQueue.push(20, 2)
	kQueue.push(25, 1)
	kQueue.push(25, 3)

	fmt.Println(kQueue.arr)

	fmt.Println("Pop 1:", kQueue.pop(1))
	fmt.Println(kQueue.arr)

	fmt.Println("Pop 3:", kQueue.pop(3))
	fmt.Println(kQueue.arr)
}
