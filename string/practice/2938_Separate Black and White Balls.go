package practice

func minimumStepsSeparateBlackWhiteBalls(s string) int64 {

	n := len(s)

	swapCount := int64(0)

	last := 0
	for cur := range n {

		if s[cur] == '0' {
			// white ball
			swapCount += int64(cur - last)
			last++
		}

	}

	return swapCount
}
