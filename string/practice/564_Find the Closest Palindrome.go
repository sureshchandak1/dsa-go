package practice

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {

	lenN := len(n)
	list := make([]int64, 0)
	isOdd := lenN%2 != 0

	mid := (lenN / 2)
	if isOdd {
		mid++
	}

	firstHalf, _ := strconv.ParseInt(n[:mid], 10, 64)

	list = append(list, findNearestPalindrome(firstHalf, isOdd))
	list = append(list, findNearestPalindrome(firstHalf+1, isOdd))
	list = append(list, findNearestPalindrome(firstHalf-1, isOdd))
	list = append(list, int64(math.Pow(10, float64(lenN-1)))-1) // all 9's
	list = append(list, int64(math.Pow(10, float64(lenN)))+1)   // 101, 1001, 10001 ...

	originalNum, _ := strconv.ParseInt(n, 10, 64)
	minDiff := int64(math.MaxInt64)
	res := int64(0)

	for _, element := range list {

		if element == originalNum {
			continue
		}

		curDiff := int64(math.Abs(float64(element - originalNum)))

		if curDiff < minDiff {
			res = element
			minDiff = curDiff
		} else if curDiff == minDiff {
			res = min(res, element)
		}
	}

	return strconv.FormatInt(res, 10)

}

func findNearestPalindrome(firstHalf int64, isOdd bool) int64 {

	res := firstHalf

	if isOdd {
		firstHalf /= 10
	}

	for firstHalf > 0 {
		res = res*10 + (firstHalf % 10)
		firstHalf /= 10
	}

	return res
}
