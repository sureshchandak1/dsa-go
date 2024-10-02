package recursion

func bubbleSort(arr []int, size int) {

	// Base case
	if size == 0 || size == 1 {
		return
	}

	isSwapped := false

	// Move large element to last
	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i] // swap value
			isSwapped = true
		}
	}

	if isSwapped {
		bubbleSort(arr, size-1)
	}
}
