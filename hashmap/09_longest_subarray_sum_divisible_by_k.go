package hashmap

func longSubarraySumDivByK(arr []int, k int) int {

	remendarMap := make(map[int]int)
	remendarMap[0] = -1 // to handle subarray starting from index 0

	sum := 0
	maxLen := 0

	for index, val := range arr {
		sum += val

		rem := sum % k

		// handle negative integer (remendar)
		if rem < 0 {
			rem += k
		}

		value, exits := remendarMap[rem]
		if exits {
			len := index - value
			maxLen = max(maxLen, len)
		} else {
			remendarMap[rem] = index
		}
	}

	return maxLen

}
