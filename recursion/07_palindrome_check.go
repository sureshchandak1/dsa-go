package recursion

func isPalindromeString(str string, start, end int) bool {

	// base case
	if start > end {
		return true
	}

	if str[start] != str[end] {
		return false
	} else {
		remaingCheck := isPalindromeString(str, start+1, end-1)
		return remaingCheck
	}
}
