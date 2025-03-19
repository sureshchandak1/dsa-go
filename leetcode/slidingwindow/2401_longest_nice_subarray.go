package slidingwindow

func longestNiceSubarray(nums []int) int {

	n := len(nums)
	start := 0
	maxLen := 0
	bitMask := 0

	for end := 0; end < n; end++ {
		// Shrinking
		for (bitMask & nums[end]) != 0 {
			bitMask = bitMask ^ nums[start]
			start++
		}

		// Expantion
		bitMask = bitMask | nums[end]

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
