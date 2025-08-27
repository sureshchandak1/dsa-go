package string

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
}

func TestPrintSubStrings(t *testing.T) {
	printSubStrings("abc")
}

func TestReverseString(t *testing.T) {
	fmt.Println(reverseString("ABCD"))
	fmt.Println(reverseString("Suresh"))
	fmt.Println(reverseString1("ABCD"))
	fmt.Println(reverseString1("Suresh"))
}

func TestIsPalindromeString(t *testing.T) {
	fmt.Println(isPalindromeString("NAMAN"))
	fmt.Println(isPalindromeString("SURESH"))
}

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
	fmt.Println(reverseWordsOptimized("the sky is blue"))
	fmt.Println(reverseWordsOptimized("  hello world  "))
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("abcd", "dacb"))
	fmt.Println(isAnagram("abcda", "dacba"))
	fmt.Println(isAnagram("abcaa", "dacba"))
}

func TestReverseWords3(t *testing.T) {
	fmt.Println(reverseWords3("We are Coders"))
}
