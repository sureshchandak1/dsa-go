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

	fmt.Println("Array Sum:", arraySum([]int{10, 20, 30, 40}, 0))
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

func TestBinarySearch(t *testing.T) {
	even := []int{2, 4, 6, 8, 12, 18}
	odd := []int{3, 8, 11, 14, 16}

	fmt.Println("isSorted:", isSorted(even, 0, len(even)))
	fmt.Println("isSorted:", isSorted(odd, 0, len(odd)))

	fmt.Println(binarySearch(even, 0, len(even)-1, 12))
	fmt.Println(binarySearch(odd, 0, len(odd)-1, 12))

	fmt.Println(binarySearch(even, 0, len(even)-1, 14))
	fmt.Println(binarySearch(odd, 0, len(odd)-1, 14))

	fmt.Println(linearSearch(even, 12, 0))
	fmt.Println(linearSearch(odd, 12, 0))
}

func TestBubbleSort(t *testing.T) {
	fmt.Println("----------Bubble Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	bubbleSort(arr, len(arr))
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	bubbleSort(arr, len(arr))
	fmt.Println(arr)
}
