package recursion

import "fmt"

// Tail recursion
func printCountingHighToLow(n int) {

	// Base case
	if n == 0 {
		return
	}

	// Processing
	fmt.Printf(" %d", n)

	// Recursion relation
	printCountingHighToLow(n - 1)
}

// Head recursion
func printCountingLowToHigh(n int) {

	// Base case
	if n == 0 {
		return
	}

	// Recursion relation
	printCountingLowToHigh(n - 1)

	// Processing
	fmt.Printf(" %d", n)
}

func factorial(n int) int {

	// Base case
	if n == 0 {
		return 1
	}

	next := factorial(n - 1)
	ans := n * next

	return ans
}

func power(n int, p int) int {

	// Base case
	if p == 0 {
		return 1
	}

	next := power(n, p-1)
	ans := n * next

	return ans
}

func arraySum(arr []int, index int) int {

	// Base case
	if index == len(arr) {
		return 0
	}

	// Recursion
	remainingPart := arraySum(arr, index+1)

	// Processing
	sum := arr[index] + remainingPart

	return sum
}
