package sorting

func mergeSort(arr []int, start, end int) {

	// Base case
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	// left part sort
	mergeSort(arr, start, mid)

	// right part sort
	mergeSort(arr, mid+1, end)

	// merge
	merge(arr, start, end)

}

func merge(arr []int, start, end int) {

	mid := start + (end-start)/2

	len1 := mid - start + 1
	len2 := end - mid

	first := make([]int, len1)
	second := make([]int, len2)

	// Copy values
	mainArrayIndex := start
	for i := 0; i < len1; i++ {
		first[i] = arr[mainArrayIndex]
		mainArrayIndex++
	}

	mainArrayIndex = mid + 1
	for i := 0; i < len2; i++ {
		second[i] = arr[mainArrayIndex]
		mainArrayIndex++
	}

	// merge 2 sorted arrays
	index1 := 0
	index2 := 0
	mainArrayIndex = start
	for index1 < len1 && index2 < len2 {
		if first[index1] < second[index2] {
			arr[mainArrayIndex] = first[index1]
			index1++
		} else {
			arr[mainArrayIndex] = second[index2]
			index2++
		}

		mainArrayIndex++
	}

	for index1 < len1 {
		arr[mainArrayIndex] = first[index1]
		index1++
		mainArrayIndex++
	}

	for index2 < len2 {
		arr[mainArrayIndex] = second[index2]
		index2++
		mainArrayIndex++
	}
}
