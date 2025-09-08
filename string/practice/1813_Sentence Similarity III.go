package practice

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {

	if len(sentence1) > len(sentence2) {
		return areSentencesSimilar(sentence2, sentence1)
	}

	smallerWords := strings.Split(sentence1, " ")
	largerWords := strings.Split(sentence2, " ")

	start := 0
	end1 := len(smallerWords) - 1
	end2 := len(largerWords) - 1

	// find prefix words
	for start <= end1 && smallerWords[start] == largerWords[start] {
		start++
	}

	// find suffix words
	for start <= end1 && smallerWords[end1] == largerWords[end2] {
		end1--
		end2--
	}

	return start > end1
}
