package sorting

func selectionSort(arr []int) {

	size := len(arr)

	for i := 0; i < size-1; i++ {

		minIndex := i
		for j := i + 1; j < size; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		// Swap 2 values
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
