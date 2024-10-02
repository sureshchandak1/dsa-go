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

func TestPalindrome(t *testing.T) {
	str := "abcde"
	fmt.Println(isPalindromeString(str, 0, len(str)-1))
	str = "abccba"
	fmt.Println(isPalindromeString(str, 0, len(str)-1))
}

func TestPower(t *testing.T) {
	fmt.Println(powerOptimized(2, 4))
	fmt.Println(powerOptimized(2, 10))
	fmt.Println(powerOptimized(3, 10))
	fmt.Println(powerOptimized(3, 11))
}

func TestReverseString(t *testing.T) {
	str := "123456789"
	runes := []rune(str)

	fmt.Println(reverseString(str))
	reverseStr(runes, 0, len(runes)-1)
	fmt.Println(string(runes))

}

func TestMergeSort(t *testing.T) {
	fmt.Println("----------Merge Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{3, 7, 0, 1, 5, 8, 3, 2, 34, 66, 87, 23, 12, 12, 12}
	fmt.Println(arr)

	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	fmt.Println("----------Quick Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	arr = []int{3, 7, 0, 1, 5, 8, 3, 2, 34, 66, 87, 23, 12, 12, 12}
	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func TestSubsequences(t *testing.T) {
	fmt.Println(getStrSubsequences("abc"))
	fmt.Println(getStrSubsequences("bbb"))
}

func TestSubsets(t *testing.T) {
	fmt.Println(getSubsets([]int{1, 2, 3}))
}
