package practice

func shiftingLetters(s string, shifts [][]int) string {

	n := len(s)
	arr := make([]int, n)

	for _, shift := range shifts {
		if shift[2] == 1 {
			// Forward direction
			arr[shift[0]]++
			if shift[1]+1 < n {
				arr[shift[1]+1]--
			}
		} else {
			// Backward direction
			arr[shift[0]]--
			if shift[1]+1 < n {
				arr[shift[1]+1]++
			}
		}
	}

	result := []rune(s)

	sum := 0
	for i := range n {
		sum = (sum + arr[i]) % 26
		if sum < 0 {
			sum += 26
		}

		shiftedChar := rune('a' + ((s[i] - 'a' + byte(sum)) % 26))
		result[i] = shiftedChar
	}

	return string(result)
}
