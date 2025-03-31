package hashmap

func findKSumSubarray(arr []int, k int) (int, int) {

	myMap := make(map[int]int)
	myMap[0] = -1

	sum := 0

	for index, val := range arr {
		sum += val

		rem := sum - k

		value, exits := myMap[rem]
		if exits {
			return value + 1, index
		}

		myMap[sum] = index
	}

	return -1, -1
}
