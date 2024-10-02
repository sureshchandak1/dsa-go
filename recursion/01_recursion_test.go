package recursion

import (
	"fmt"
	"testing"
)

func TestRecursion(t *testing.T) {
	printCountingHighToLow(5)
	fmt.Println()
	printCountingLowToHigh(5)
	fmt.Println()

	// Factorial 4 = 4 * 3 * 2 * 1
	fmt.Println("Factorial:", factorial(4))
	fmt.Println("Factorial:", factorial(5))
	fmt.Println("Factorial:", factorial(6))

	fmt.Println("Power:", power(2, 4))
}

func TestClimbStair(t *testing.T) {
	reachHome(1, 5)
	fmt.Println(climbStairWays(4))
	fmt.Println(climbStairWays(6))
}

func TestFibonacci(t *testing.T) {
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(10))
}

func TestSatDigits(t *testing.T) {
	counts := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	sayDigits(412, counts)
}
