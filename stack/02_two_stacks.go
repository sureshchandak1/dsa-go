package stack

import "fmt"

type twoStack struct {
	size  int
	stack []int
	top1  int
	top2  int
}

func (s *twoStack) push1(data int) {
	if s.top2-s.top1 > 1 {
		s.top1++
		s.stack[s.top1] = data
	}
}

func (s *twoStack) push2(data int) {
	if s.top2-s.top1 > 1 {
		s.top2--
		s.stack[s.top2] = data
	}
}

func (s *twoStack) pop1() int {
	if s.top1 >= 0 {
		data := s.stack[s.top1]
		s.stack[s.top1] = 0
		s.top1--
		return data
	}
	return -1
}

func (s *twoStack) pop2() int {
	if s.top2 < s.size {
		data := s.stack[s.top2]
		s.stack[s.top2] = 0
		s.top2++
		return data
	}
	return -1
}

func (s *twoStack) print() {
	fmt.Print("[ ")
	for i := 0; i < s.size; i++ {
		fmt.Print(s.stack[i])
		fmt.Print(" ")
	}
	fmt.Print("]")
	fmt.Println()
}
