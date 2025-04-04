package hashmap

func largestSubarray0and1(arr []int) int {

	sumMap := make(map[int]int)
	sumMap[0] = -1 // to handle subarray with sum 0 and starting from index 0

	len, maxLen := 0, 0
	sum := 0

	for index, val := range arr {
		if val == 0 {
			sum -= 1
		} else {
			sum += 1
		}

		_, exits := sumMap[sum]
		if exits {
			len = index - sumMap[sum]
			maxLen = max(maxLen, len)
		} else {
			sumMap[sum] = index
		}

	}

	return maxLen
}
