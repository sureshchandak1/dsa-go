package practice

import (
	"fmt"
	"testing"
)

func TestCountConsistentStrings(t *testing.T) {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}

func TestLengthOfLastWord(t *testing.T) {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func TestMakeGoodString(t *testing.T) {
	fmt.Println(makeGoodString("leEeetcode"))
	fmt.Println(makeGoodString("abBAcC"))
	fmt.Println(makeGoodString("s"))
	fmt.Println(makeGoodStringStack("leEeetcode"))
	fmt.Println(makeGoodStringStack("abBAcC"))
	fmt.Println(makeGoodStringStack("s"))
}
