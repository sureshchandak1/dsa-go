package hashmap

func findLongestKSumSubarray(arr []int, k int) int {

	maxLen := 0

	myMap := make(map[int]int)
	myMap[0] = -1

	sum := 0

	for index, val := range arr {
		sum += val

		rem := sum - k

		value, exits := myMap[rem]
		if exits {
			len := index - value
			maxLen = max(maxLen, len)
		}

		_, ok := myMap[sum]
		if !ok {
			myMap[sum] = index
		}
	}

	return maxLen
}
