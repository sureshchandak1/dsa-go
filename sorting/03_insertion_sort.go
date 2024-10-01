package sorting

func insertionSort(arr []int) {

	size := len(arr)

	for i := 1; i < size; i++ {

		temp := arr[i]

		// swift large values
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > temp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		// insert
		arr[j+1] = temp
	}
}
