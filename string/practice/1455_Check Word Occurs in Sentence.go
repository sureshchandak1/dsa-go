package practice

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {

	n := len(sentence)
	m := len(searchWord)
	wordCount := 1

	i := 0
	j := 0
	for i < n {

		for i < n && sentence[i] == ' ' {
			i++
		}

		// compare the words
		for i < n && j < m && sentence[i] == searchWord[j] {
			i++
			j++
		}

		if j == m {
			return wordCount
		}

		j = 0

		// move to next space
		for i < n && sentence[i] != ' ' {
			i++
		}

		// move to next word
		i++
		wordCount++
	}

	return -1
}

func isPrefixOfWordSolve(sentence string, searchWord string) int {

	words := strings.Split(sentence, " ")

	m := len(searchWord)
	n := len(words)

	for i := 0; i < n; i++ {
		if len(words[i]) >= m {
			sub := words[i][:m]
			if sub == searchWord {
				return i + 1
			}
		}
	}

	return -1
}
