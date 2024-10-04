package stack

import "fmt"

type stack struct {
	size  int
	stack []int
	top   int
}

func (s *stack) push(data int) {
	if s.size-s.top > 1 {
		s.top++
		s.stack[s.top] = data
	} else {
		fmt.Println("Stack OverFlow")
	}
}

func (s *stack) pop() {
	if s.top >= 0 {
		s.top--
	} else {
		fmt.Println("Stack Empty")
	}
}

func (s *stack) peek() int {
	if s.top >= 0 {
		return s.stack[s.top]
	} else {
		fmt.Println("Stack Empty")
		return -1
	}
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) length() int {
	return s.top + 1
}

func (s *stack) print() {
	fmt.Print("[ ")
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.stack[i])
		fmt.Print(" ")
	}
	fmt.Print("]")
	fmt.Println()
}
