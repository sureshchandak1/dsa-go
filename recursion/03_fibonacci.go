package recursion

func fibonacci(n int) int {

	// Base case
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	value1 := fibonacci(n - 1)
	value2 := fibonacci(n - 2)

	sum := value1 + value2

	return sum
}
