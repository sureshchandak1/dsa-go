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

func TestTwoStack(t *testing.T) {

	s := twoStack{
		size:  10,
		stack: make([]int, 10),
		top1:  -1,
		top2:  10,
	}

	s.push1(11)
	s.push1(12)
	s.push2(13)
	s.push2(14)
	s.push1(15)
	s.push2(16)
	s.print()
	fmt.Println("Pop1:", s.pop1())
	fmt.Println("Pop2:", s.pop2())
	s.print()
}
