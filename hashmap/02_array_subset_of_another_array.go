package hashmap

func isSubset(arr1 []int, arr2 []int) bool {

	freqMap := make(map[int]int)

	for _, val := range arr1 {
		freqMap[val]++
	}

	for _, val := range arr2 {
		count, exits := freqMap[val]

		if exits && count > 0 {
			freqMap[val]--
		} else {
			return false
		}
	}

	return true
}
