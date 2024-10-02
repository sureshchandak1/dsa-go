package recursion

func quickSort(arr []int, start, end int) {

	// Base case
	if start >= end {
		return
	}

	// partition
	p := partition(arr, start, end)

	// left part sorting
	quickSort(arr, start, p-1)

	// right part sorting
	quickSort(arr, p+1, end)

}

func partition(arr []int, start, end int) int {

	pivot := arr[start]

	count := 0
	for i := start + 1; i <= end; i++ {
		if arr[i] <= pivot {
			count++
		}
	}

	// place pivot at right position
	pivotIndex := start + count
	arr[pivotIndex], arr[start] = arr[start], arr[pivotIndex] // swap values

	// handle pivot left and right part
	i := start
	j := end

	for i < pivotIndex && j > pivotIndex {

		for arr[i] <= pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		// swap values
		if i < pivotIndex && j > pivotIndex {
			arr[i], arr[j] = arr[j], arr[i]
		}

		i++
		j--
	}

	return pivotIndex
}
