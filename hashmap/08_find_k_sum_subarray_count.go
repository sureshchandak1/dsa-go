package hashmap

func findKSumSubarrayCount(arr []int, k int) int {

	sumFreqMap := make(map[int]int)
	sumFreqMap[0] = 1

	sum := 0

	count := 0

	for _, val := range arr {
		sum += val

		rem := sum - k

		value, exits := sumFreqMap[rem]
		if exits {
			count += value
		}

		sumFreqMap[sum]++
	}

	return count
}
