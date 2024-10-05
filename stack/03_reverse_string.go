package stack

import (
	"container/list"
)

func reverseString(str string) string {

	runes := []rune(str)
	stack := list.New()

	for i := 0; i < len(runes); i++ {
		stack.PushBack(runes[i])
	}

	revStr := []rune{}
	for stack.Len() != 0 {
		back := stack.Back()
		revStr = append(revStr, back.Value.(int32))
		stack.Remove(back)
	}

	return string(revStr)

}
