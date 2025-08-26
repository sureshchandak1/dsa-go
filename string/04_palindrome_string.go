package string

func isPalindromeString(str string) bool {

	start := 0
	end := len(str) - 1

	for start < end {

		if str[start] != str[end] {
			return false
		}

		start++
		end--
	}

	return true
}
