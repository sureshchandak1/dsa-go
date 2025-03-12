package slidingwindow

func maxSumSubarray(arr []int, k int) int {

	n := len(arr)
	sum := 0

	// solve first window
	index := 0
	for index < n && index < k { // T.C - K
		sum += arr[index]
		index++
	}

	maxSum := sum

	prevElement, nextElement := 0, 0
	for index := 1; index < n-k+1; index++ { // T.C - N-K
		prevElement = arr[index-1]
		nextElement = arr[index+k-1]

		sum = sum - prevElement + nextElement

		maxSum = max(maxSum, sum)
	}

	// T.C: K + N - K => N
	// S.C: Constant O(1)
	return maxSum
}
