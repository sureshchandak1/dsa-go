package practice

import (
	"strings"
	"unicode"
)

func clearDigits(s string) string {

	var result strings.Builder

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			result.WriteRune(ch)
		} else {
			str := result.String()
			if len(str) > 0 {
				result.Reset()
				result.WriteString(str[:len(str)-1])
			}
		}
	}

	return result.String()
}

func clearDigitsWithStack(s string) string {
	stack := []rune{}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	var result strings.Builder
	for i := len(stack) - 1; i >= 0; i-- {
		result.WriteRune(stack[i])
	}

	// Reverse the string because we inserted from the end
	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
