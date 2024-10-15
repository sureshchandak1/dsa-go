package stack

import (
	"fmt"
	"testing"
)

func TestMinStack(t *testing.T) {

	minStack := createMinStack()
	minStack.push(5)
	minStack.push(3)
	minStack.push(7)
	minStack.push(2)

	fmt.Println("Top:", minStack.top())
	fmt.Println("Min:", minStack.getMin())

	minStack.pop()

	fmt.Println("Top:", minStack.top())
	fmt.Println("Min:", minStack.getMin())

	minStack.pop()

	fmt.Println("Top:", minStack.top())
	fmt.Println("Min:", minStack.getMin())

}
