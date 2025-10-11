package practice

import (
	"math"
	"slices"
)

func findRotateSteps(ring string, key string) int {

	ringLen := len(ring)
	keyLen := len(key)

	posArr := make([][]int, 26)

	for i := range ringLen {
		index := ring[i] - 'a'
		posArr[index] = append(posArr[index], i)
	}

	dp := make([][]int, keyLen)
	for i := range dp {
		dp[i] = slices.Repeat([]int{-1}, ringLen)
	}

	return getMinSteps(key, ring, 0, 0, posArr, dp)
}

func getMinSteps(key, ring string, ringIndex, keyIndex int, posArr [][]int, dp [][]int) int {

	// Base index
	if keyIndex == len(key) {
		return 0
	}

	if dp[keyIndex][ringIndex] != -1 {
		return dp[keyIndex][ringIndex]
	}

	// generate all possibilities
	charIndex := key[keyIndex] - 'a'
	positions := posArr[charIndex]

	minAns := math.MaxInt32
	ringLen := len(ring)

	for _, pos := range positions {

		clockwiseStep := abs(ringIndex - pos)
		antiClockwiseStep := ringLen - clockwiseStep
		minStep := min(clockwiseStep, antiClockwiseStep)

		curAns := minStep + getMinSteps(key, ring, pos, keyIndex+1, posArr, dp)
		if curAns < minAns {
			minAns = curAns
		}
	}

	dp[keyIndex][ringIndex] = minAns + 1 // 1 is for the button press

	return dp[keyIndex][ringIndex]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
