package sorting

import (
	"fmt"
	"testing"
)

func TestSortPackage(t *testing.T) {
	arr := []int{10, 1, 4, 6, 12, 50, 70, 0, -1, 4}
	fmt.Println(arr)
	sortIntArray(arr, true)
	fmt.Println(arr)
	sortIntArray(arr, false)
	fmt.Println(arr)
}

func TestSort2dArray(t *testing.T) {
	matrix := [][]int{
		{2, 3, 1},
		{6, 3, 5},
		{1, 4, 9},
	}
	fmt.Println(matrix)
	sort2dArray(matrix)
	fmt.Println(matrix)
}

func TestSelectionSort(t *testing.T) {
	fmt.Println("----------Selection Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	selectionSort(arr)
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	selectionSort(arr)
	fmt.Println(arr)
}

func TestBubbleSort(t *testing.T) {
	fmt.Println("----------Bubble Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	bubbleSort(arr)
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	bubbleSort(arr)
	fmt.Println(arr)
}

func TestInsertionSort(t *testing.T) {
	fmt.Println("----------Insertion Sort----------")

	arr := []int{29, 72, 98, 13, 87, 66, 52, 51, 36}
	fmt.Println(arr)

	insertionSort(arr)
	fmt.Println(arr)

	arr = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(arr)

	insertionSort(arr)
	fmt.Println(arr)
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
