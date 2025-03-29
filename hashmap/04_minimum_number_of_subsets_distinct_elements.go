package hashmap

// Brout force
func findMinNumberOfSets1(arr []int) int {
	res := 0

	n := len(arr)
	visited := make([]bool, n)

	for i := range n {
		if visited[i] {
			continue
		}

		res++

		set := make(map[int]bool)

		for j := i; j < n; j++ {
			if visited[j] || set[arr[j]] {
				continue
			}

			visited[j] = true
			set[arr[j]] = true
		}
	}

	return res
}

// Optimized
func findMinNumberOfSets2(arr []int) int {

	freqMap := make(map[int]int)

	maxFreq := 0

	for _, val := range arr {
		freqMap[val]++

		maxFreq = max(maxFreq, freqMap[val])
	}

	return maxFreq
}
