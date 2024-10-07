package stack

import (
	"container/list"
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

func TestReverseString(t *testing.T) {
	fmt.Println(reverseString("abc"))
	fmt.Println(reverseString("love"))
}

func TestDeleteMiddleElement(t *testing.T) {
	stack := list.New()
	stack.PushBack(10)
	stack.PushBack(20)
	stack.PushBack(30)
	stack.PushBack(40)
	stack.PushBack(50)
	stack.PushBack(60)
	stack.PushBack(70)

	deleteMiddleElement(stack)

	printStack(stack)

}

func TestIsValidParenthesis(t *testing.T) {
	fmt.Println(isValidParenthesis("[()]{}{[()()]()}"))
	fmt.Println(isValidParenthesis("[[}["))
	fmt.Println(isValidParenthesis("[(])"))
}

func TestPushBottom(t *testing.T) {
	stack := list.New()
	stack.PushBack(10)
	stack.PushBack(20)
	stack.PushBack(30)
	stack.PushBack(40)
	stack.PushBack(50)
	stack.PushBack(60)
	stack.PushBack(70)

	fmt.Println("Stack Top:", stack.Back().Value)

	insertStackBottom(stack, 80)

	fmt.Println("Stack Top:", stack.Back().Value)

	printStack(stack)

}

func TestReverseStack(t *testing.T) {
	stack := list.New()
	stack.PushBack(10)
	stack.PushBack(20)
	stack.PushBack(30)
	stack.PushBack(40)
	stack.PushBack(50)
	stack.PushBack(60)
	stack.PushBack(70)

	fmt.Println("Stack Top:", stack.Back().Value)

	reverseStack(stack)

	fmt.Println("Stack Top:", stack.Back().Value)

	printStack(stack)

}

func TestSortStack(t *testing.T) {
	stack := list.New()
	stack.PushBack(70)
	stack.PushBack(60)
	stack.PushBack(50)
	stack.PushBack(40)
	stack.PushBack(30)
	stack.PushBack(20)
	stack.PushBack(10)

	fmt.Println("Stack Top:", stack.Back().Value)

	sortStack(stack)

	fmt.Println("Stack Top:", stack.Back().Value)

	printStack(stack)

}

func printStack(stack *list.List) {
	fmt.Print("[ ")
	for stack.Len() != 0 {
		back := stack.Back()

		fmt.Print(back.Value)
		fmt.Print(" ")

		stack.Remove(back)
	}
	fmt.Print("]")
	fmt.Println()
}
