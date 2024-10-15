package stack

import "container/list"

type minStack struct {
	stack *list.List
	min   int
}

func createMinStack() minStack {
	minStack := minStack{}

	minStack.stack = list.New()

	return minStack
}

func (m *minStack) push(data int) {
	if m.stack.Len() == 0 {
		m.stack.PushBack(data)
		m.min = data
	} else {
		if data < m.min {
			m.stack.PushBack(2*data - m.min)
			m.min = data
		} else {
			m.stack.PushBack(data)
		}
	}
}

func (m *minStack) pop() {
	if m.stack.Len() == 0 {
		return
	}

	curr := m.stack.Back()
	currVal := curr.Value.(int)
	m.stack.Remove(curr)

	if currVal > m.min {
		return
	} else {
		val := 2*m.min - currVal
		m.min = val
	}
}

func (m *minStack) top() int {

	if m.stack.Len() == 0 {
		return -1
	}

	curr := m.stack.Back()
	currVal := curr.Value.(int)

	if currVal < m.min {
		return m.min
	} else {
		return currVal
	}
}

func (m *minStack) getMin() int {

	if m.stack.Len() == 0 {
		return -1
	}

	return m.min
}
