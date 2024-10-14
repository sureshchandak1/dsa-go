package stack

import (
	"fmt"
	"testing"
)

func TestNStack(t *testing.T) {
	nStack := createNStack(3, 6)
	nStack.push(1, 10)
	nStack.push(1, 20)
	nStack.push(2, 30)
	nStack.push(3, 40)
	nStack.push(3, 50)

	fmt.Println(nStack.arr)

	fmt.Println("Pop from stack 1:", nStack.pop(1))

	fmt.Println(nStack.arr)

	nStack.push(1, 60)
	fmt.Println(nStack.arr)

	fmt.Println("Pop from stack 2:", nStack.pop(2))
	fmt.Println(nStack.arr)

	nStack.push(1, 70)
	fmt.Println(nStack.arr)
}
