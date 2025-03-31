package hashmap

func checkSumSubarray(arr []int, k int) bool {
	set := make(map[int]bool)
	set[0] = true

	sum := 0

	for _, val := range arr {
		sum += val

		rem := sum - k

		if set[rem] {
			return true
		}

		set[sum] = true
	}

	return false
}
