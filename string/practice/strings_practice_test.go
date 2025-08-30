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

func TestReversePrefix(t *testing.T) {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}

func TestCompareVersion(t *testing.T) {
	fmt.Println(compareVersion("1.2", "1.10"))
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0.0"))
	fmt.Println(compareVersionOptimized("1.2", "1.10"))
	fmt.Println(compareVersionOptimized("1.01", "1.001"))
	fmt.Println(compareVersionOptimized("1.0", "1.0.0.0"))
}

func TestFindRelativeRanks(t *testing.T) {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}

func TestScoreOfString(t *testing.T) {
	fmt.Println(scoreOfString("hello"))
	fmt.Println(scoreOfString("zaz"))
}

func TestReverseByteArray(t *testing.T) {
	reverseByteArray([]rune{'h', 'e', 'l', 'l', 'o'})
}

func TestAppendCharacters(t *testing.T) {
	fmt.Println(appendCharacters("coaching", "coding"))
	fmt.Println(appendCharacters("abcde", "a"))
	fmt.Println(appendCharacters("z", "abcde"))
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindromeOptimized("abccccdd"))
	fmt.Println(longestPalindromeOptimized("a"))
}

func TestCommonChars(t *testing.T) {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))
}

func TestIsNStraightHand(t *testing.T) {
	fmt.Println(isNStraightHand([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(isNStraightHand([]int{1, 2, 3, 4, 5}, 4))
}

func TestReplaceWords(t *testing.T) {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
	fmt.Println(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"))
}
