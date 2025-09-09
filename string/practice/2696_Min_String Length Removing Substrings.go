package practice

func minLengthRemoveSubstringOptimized(s string) int {

	str := []rune(s)
	n := len(str)

	write := 0
	for read := range n {

		if write == 0 {
			str[write] = str[read]
			write++
			continue
		}

		if str[read] == 'B' && str[write-1] == 'A' {
			write--
		} else if str[read] == 'D' && str[write-1] == 'C' {
			write--
		} else {
			str[write] = str[read]
			write++
		}

	}

	return write
}

func minLengthRemoveSubstring(s string) int {

	stack := []rune{}

	for _, ch := range s {

		if len(stack) == 0 {
			stack = append(stack, ch)
			continue
		}

		if ch == 'B' && stack[len(stack)-1] == 'A' {
			stack = stack[:len(stack)-1]
		} else if ch == 'D' && stack[len(stack)-1] == 'C' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack)
}
