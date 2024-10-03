package recursion

import "sort"

func uniquePermutations(str string) []string {
	result := []string{}
	premMap := make(map[string]bool)

	if isAllSameChar(str) {
		result = append(result, str)
	} else {

		permutations([]rune(str), 0, premMap)

		for key := range premMap {
			result = append(result, key)
		}

		sort.Strings(result)
	}

	return result
}

func permutations(str []rune, index int, premMap map[string]bool) {

	// Base case
	if index >= len(str) {
		premMap[string(str)] = true
		return
	}

	for j := index; j < len(str); j++ {
		str[index], str[j] = str[j], str[index] // swap values

		permutations(str, index+1, premMap)

		// backtrack (swap value in same string so back to original string after work done)
		str[index], str[j] = str[j], str[index] // back to previous value
	}
}

func isAllSameChar(str string) bool {

	for i := 1; i < len(str); i++ {
		if str[i] != str[0] {
			return false
		}
	}

	return true
}

func uniquePermutationsArray(arr []int) [][]int {
	result := [][]int{}
	permutationsArray(arr, 0, &result)
	return result
}

func permutationsArray(arr []int, index int, result *[][]int) {

	// Base case
	if index >= len(arr) {
		*result = append(*result, append(make([]int, 0), arr...))
		return
	}

	for j := index; j < len(arr); j++ {

		arr[index], arr[j] = arr[j], arr[index] // swap value

		permutationsArray(arr, index+1, result)

		// Backtrack to previous arr
		arr[index], arr[j] = arr[j], arr[index] // swap values

	}
}
