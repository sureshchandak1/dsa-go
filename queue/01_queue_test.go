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

func TestCircularQueue(t *testing.T) {
	queue := newCircularQueue(5)
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

	fmt.Println("First:", queue.first())

	fmt.Println(queue.arr)

	fmt.Println("Pop:", queue.pop())
	fmt.Println("First:", queue.first())

	fmt.Println(queue.arr)
}

func TestDeque(t *testing.T) {
	queue := newDeque(10)
	queue.pushRear(10)
	queue.pushRear(20)
	queue.pushRear(30)
	queue.pushRear(40)
	queue.pushRear(50)
	queue.pushFront(60)
	queue.pushFront(70)
	queue.pushFront(80)
	queue.pushFront(90)
	queue.pushFront(100)

	fmt.Println(queue.arr)

	fmt.Println("Front:", queue.getFront())
	fmt.Println("Rear:", queue.getRear())

	fmt.Println("Pop Front:", queue.popFront())
	fmt.Println("Pop Rear:", queue.popRear())

	fmt.Println(queue.arr)

	fmt.Println("Pop Rear:", queue.popRear())
	fmt.Println("Pop Rear:", queue.popRear())

	fmt.Println(queue.arr)

	queue.pushFront(101)
	queue.pushFront(102)

	fmt.Println(queue.arr)

	fmt.Println("Pop Front:", queue.popFront())
	fmt.Println("Pop Front:", queue.popFront())

	queue.pushRear(1001)
	queue.pushRear(2002)

	fmt.Println(queue.arr)
}
