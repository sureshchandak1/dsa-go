package recursion

func getStrSubsequences(str string) []string {
	result := []string{}
	subsequences(str, "", 0, &result)
	return result
}

func subsequences(str string, output string, index int, result *[]string) {

	// Base case
	if index >= len(str) {
		if len(output) > 0 {
			*result = append(*result, output)
		}
		return
	}

	// exclude
	subsequences(str, output, index+1, result)

	var currChar byte = str[index]

	// include
	subsequences(str, output+string(currChar), index+1, result)
}
