package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	s := stack{
		size:  5,
		stack: make([]int, 5),
		top:   -1,
	}

	fmt.Println("isEmpty:", s.isEmpty())
	s.push(11)
	s.push(12)
	s.push(13)
	s.push(14)
	s.push(15)
	s.push(16)
	s.print()
	fmt.Println("Size:", s.length())
	fmt.Println("Peek:", s.peek())
	s.pop()
	s.pop()
	s.print()
	fmt.Println("Size:", s.length())
	fmt.Println("Peek:", s.peek())
	fmt.Println("isEmpty:", s.isEmpty())
}
