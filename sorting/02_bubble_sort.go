package sorting

func bubbleSort(arr []int) {

	size := len(arr)

	for i := 0; i < size-1; i++ {

		isSwapped := false

		for j := 0; j < (size-i)-1; j++ {
			if arr[j] > arr[j+1] {
				// swap values
				arr[j], arr[j+1] = arr[j+1], arr[j]

				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}
}
