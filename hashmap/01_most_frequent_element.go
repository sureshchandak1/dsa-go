package hashmap

func mostFrequentElement(arr []int) int {

	freqMap := make(map[int]int)

	ans := -1
	maxFreq := 0

	for _, val := range arr {

		freqMap[val]++

		if freqMap[val] > maxFreq {
			maxFreq = freqMap[val]
			ans = val
		}
	}

	return ans
}
