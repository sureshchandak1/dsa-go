package slidingwindow

func smallestSubarrayGreaterX(arr []int, x int) int {

	n := len(arr)
	windowStart := 0
	windowsEnd := 0
	sum := 0
	minLen := 234234213
	len := 0

	for windowsEnd < n {

		// expansion
		sum += arr[windowsEnd]

		if sum > x {
			len = windowsEnd - windowStart + 1
			minLen = min(minLen, len)

			// Shrinking
			for windowStart < windowsEnd && sum > x {
				sum -= arr[windowStart]
				windowStart++
				// if sum greater then x then may be posible answer
				if sum > x {
					len = windowsEnd - windowStart + 1
					minLen = min(minLen, len)
				}
			}
		}

		windowsEnd++
	}

	if minLen == 234234213 {
		return 0
	} else {
		return minLen
	}

}
