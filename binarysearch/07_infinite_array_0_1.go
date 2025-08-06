package binarysearch

func findFirstIndexOf1(arr []int) int {

	rangeIndex := findRangeInfiniteArr(arr, 1)

	return firstIndexFind(arr, 1, rangeIndex[0], rangeIndex[1])

}

func firstIndexFind(arr []int, target, start, end int) int {

	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		if arr[mid] == target {
			ans = mid
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

		mid = start + (end-start)/2
	}

	return ans
}
