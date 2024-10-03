package recursion

var keypad = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	result := []string{}
	createCombinations(digits, "", 0, &result)
	return result
}

func createCombinations(digits string, combination string, index int, result *[]string) {

	// base case
	if index >= len(digits) {
		*result = append(*result, combination)
		return
	}

	var digit byte = digits[index]
	mapping := keypad[digit-'0']

	for _, ch := range mapping {
		createCombinations(digits, combination+string(ch), index+1, result)
	}
}
