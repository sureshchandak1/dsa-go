package string

import (
	"bytes"
	"strings"
)

func reverseWords(str string) string {

	arr := strings.Split(str, " ")

	buffer := bytes.NewBufferString("")

	for i := len(arr) - 1; i >= 0; i-- {

		if len(arr[i]) > 0 {

			if len(buffer.String()) == 0 {
				buffer.WriteString(arr[i])
			} else {
				buffer.WriteString(" ")
				buffer.WriteString(arr[i])
			}

		}
	}

	return buffer.String()
}

func reverseWordsOptimized(str string) string {

	var buffer bytes.Buffer

	start := len(str) - 1

	for start >= 0 {

		for start >= 0 && str[start] == ' ' {
			start--
		}

		if start < 0 {
			break
		}

		endIndex := start

		for start >= 0 && str[start] != ' ' {
			start--
		}

		if len(buffer.String()) == 0 {
			buffer.WriteString(str[start+1 : endIndex+1])
		} else {
			buffer.WriteString(" ")
			buffer.WriteString(str[start+1 : endIndex+1])
		}
	}

	return buffer.String()
}
