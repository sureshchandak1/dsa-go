package leetcode

func minimumRecolors(blocks string, k int) int {

	wCount := 0

	// First window
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wCount++
		}
	}

	result := wCount
	n := len(blocks)

	// Next window
	for i := 1; i < (n - k + 1); i++ {
		if blocks[i-1] == 'W' {
			wCount--
		}
		if blocks[i+k-1] == 'W' {
			wCount++
		}

		result = min(result, wCount)
	}

	return result

}
