package stack

type nStack struct {
	arr      []int
	top      []int
	next     []int
	freeSpot int
}

func createNStack(totalStack, arrSize int) nStack {
	nStack := nStack{
		arr:      make([]int, arrSize),
		top:      make([]int, totalStack),
		next:     make([]int, arrSize),
		freeSpot: 0,
	}

	// insert default values
	// top initialise
	for i := 0; i < totalStack; i++ {
		nStack.top[i] = -1
	}

	// next initialise
	for i := 0; i < arrSize; i++ {
		nStack.next[i] = i + 1
	}
	// update last index
	nStack.next[arrSize-1] = -1

	return nStack
}

func (s *nStack) push(stack, data int) bool {
	// Check for overflow
	if s.freeSpot == -1 {
		return false
	}

	// find index
	index := s.freeSpot

	// insert element into array
	s.arr[index] = data

	// update freeSpot
	s.freeSpot = s.next[index]

	// update next
	s.next[index] = s.top[stack-1]

	// update top
	s.top[stack-1] = index

	return true
}

func (s *nStack) pop(stack int) int {

	if s.top[stack-1] == -1 {
		return -1
	}

	index := s.top[stack-1]

	s.top[stack-1] = s.next[index]

	s.next[index] = s.freeSpot

	s.freeSpot = index

	value := s.arr[index]
	s.arr[index] = 0

	return value
}
