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

func TestReverseParenthesesString(t *testing.T) {
	fmt.Println(reverseParenthesesString("(abcd)"))
	fmt.Println(reverseParenthesesString("(u(love)i)"))
	fmt.Println(reverseParenthesesString("(ed(et(oc))el)"))
	fmt.Println(reverseParenthesesStringOptimized("(abcd)"))
	fmt.Println(reverseParenthesesStringOptimized("(u(love)i)"))
	fmt.Println(reverseParenthesesStringOptimized("(ed(et(oc))el)"))
}

func TestMaximumGain(t *testing.T) {
	fmt.Println(maximumGain("cdbcbbaaabab", 4, 5))
	fmt.Println(maximumGain("aabbaaxybbaabb", 5, 4))
}

func TestCountSeniors(t *testing.T) {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(countSeniors([]string{"1313579440F2036", "2921522980M5644"}))
	fmt.Println(countSeniorsOptimized([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(countSeniorsOptimized([]string{"1313579440F2036", "2921522980M5644"}))
}

func TestFractionAddition(t *testing.T) {
	fmt.Println(fractionAddition("-1/2+1/2"))
	fmt.Println(fractionAddition("-1/2+1/2+1/3"))
	fmt.Println(fractionAddition("1/3-1/2"))
}

func TestNearestPalindromic(t *testing.T) {
	fmt.Println(nearestPalindromic("123"))
	fmt.Println(nearestPalindromic("1"))
}

func TestGetSumStringDigits(t *testing.T) {
	fmt.Println(getSumStringDigits("iiii", 1))
	fmt.Println(getSumStringDigits("leetcode", 2))
	fmt.Println(getSumStringDigits("zbax", 2))
}

func TestFindMinDifference(t *testing.T) {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))
}

func TestUncommonFromSentences(t *testing.T) {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
	fmt.Println(uncommonFromSentences("apple apple", "banana"))
}

func TestLargestNumber(t *testing.T) {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}

func TestPermutationInString(t *testing.T) {
	fmt.Println(permutationInString("ab", "eidbaooo"))
	fmt.Println(permutationInString("ab", "eidboaoo"))
	fmt.Println(permutationInStringOptimized("ab", "eidbaooo"))
	fmt.Println(permutationInStringOptimized("ab", "eidboaoo"))
}

func TestAreSentencesSimilar(t *testing.T) {
	fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	fmt.Println(areSentencesSimilar("of", "A lot of words"))
	fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
}

func TestMinLengthRemoveSubstring(t *testing.T) {
	fmt.Println(minLengthRemoveSubstring("ABFCACDB"))
	fmt.Println(minLengthRemoveSubstring("ACBBD"))
	fmt.Println(minLengthRemoveSubstringOptimized("ABFCACDB"))
	fmt.Println(minLengthRemoveSubstringOptimized("ACBBD"))
}

func TestMinimumStepsSeparateBlackWhiteBalls(t *testing.T) {
	fmt.Println(minimumStepsSeparateBlackWhiteBalls("101"))
	fmt.Println(minimumStepsSeparateBlackWhiteBalls("100"))
	fmt.Println(minimumStepsSeparateBlackWhiteBalls("0111"))
}

func TestLongestDiverseString(t *testing.T) {
	fmt.Println(longestDiverseString(1, 1, 7))
	fmt.Println(longestDiverseString(7, 1, 0))
}

func TestFindKthBit(t *testing.T) {
	fmt.Println(findKthBit(3, 1))
	fmt.Println(findKthBit(4, 11))
}

func TestParseBoolExpr(t *testing.T) {
	fmt.Println(parseBoolExpr("&(|(f))"))
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
	fmt.Println(parseBoolExpr("!(&(f,t))"))
}

func TestMaxUniqueSplit(t *testing.T) {
	fmt.Println(maxUniqueSplit("ababccc"))
	fmt.Println(maxUniqueSplit("aba"))
	fmt.Println(maxUniqueSplit("aa"))
}
