package hashmap

func pairCount1(arr []int, k int) int {

	freqMap := make(map[int]int)

	for _, val := range arr {
		freqMap[val]++
	}

	ans := 0

	for _, val := range arr {
		if val > k {
			continue
		}

		secondVal := k - val

		count, exits := freqMap[secondVal]
		if exits && count > 0 {
			ans += freqMap[secondVal]

			freqMap[val]--

			// ignore same value pair
			if val == secondVal {
				ans -= 1
			}
		}
	}

	return ans
}

func pairCount2(arr []int, k int) int {

	freqMap := make(map[int]int)
	ans := 0

	for _, val := range arr {
		if val > k {
			continue
		}

		ans += freqMap[k-val]
		freqMap[val]++
	}

	return ans
}
