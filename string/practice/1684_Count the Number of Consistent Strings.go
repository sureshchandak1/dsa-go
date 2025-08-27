package practice

/**
func countConsistentStrings(allowed string, words []string) int {

	counter := 0
	set := make(map[rune]struct{})

	for _, char := range allowed {
		set[char] = struct{}{}
	}

	for _, str := range words {
		flag := true

		for _, char := range str {
			if _, exists := set[char]; !exists {
				flag = false
				break
			}
		}

		if flag {
			counter++
		}
	}

	return counter
}
	**/

func countConsistentStrings(allowed string, words []string) int {
	isAllowed := make([]bool, 26)

	for _, c := range allowed {
		isAllowed[c-'a'] = true
	}

	count := 0

	for _, word := range words {
		if isConsistent(word, isAllowed) {
			count++
		}
	}

	return count
}

func isConsistent(word string, isAllowed []bool) bool {
	for i := range len(word) {
		if !isAllowed[word[i]-'a'] {
			return false
		}
	}
	return true
}
