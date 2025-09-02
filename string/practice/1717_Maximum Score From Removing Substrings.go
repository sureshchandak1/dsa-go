package practice

var str []rune

func maximumGain(sInput string, x, y int) int {

	n := len(sInput)
	str = make([]rune, n+1)
	for i, ch := range sInput {
		str[i] = ch
	}
	str[n] = '$'

	return solve(sInput, x, y)
}

func solve(_ string, x, y int) int {

	var firstPair, secondPair string
	if x > y {
		firstPair = "ab"
		secondPair = "ba"
	} else {
		firstPair = "ba"
		secondPair = "ab"
	}

	score := 0

	// score += removeSubstring(firstPair, max(x, y))
	// score += removeSubstring(secondPair, min(x, y))

	score += removeSubstringOptimized(firstPair, max(x, y))
	score += removeSubstringOptimized(secondPair, min(x, y))

	return score
}

func removeSubstringOptimized(pair string, score int) int {
	totalScore := 0

	firstTargetChar := rune(pair[0])
	secondTargetChar := rune(pair[1])

	writeIndex := 0
	for readIndex := 0; str[readIndex] != '$'; readIndex++ {
		str[writeIndex] = str[readIndex]
		writeIndex++

		if writeIndex > 1 {
			firstChar := str[writeIndex-2]
			secondChar := str[writeIndex-1]

			if firstChar == firstTargetChar && secondChar == secondTargetChar {
				writeIndex -= 2
				totalScore += score
			}
		}
	}

	str[writeIndex] = '$'

	return totalScore
}

func removeSubstring(pair string, score int) int {
	totalScore := 0

	stack := []rune{}

	firstChar := rune(pair[0])
	secondChar := rune(pair[1])

	for _, ch := range str {
		if len(stack) > 0 && ch == secondChar && stack[len(stack)-1] == firstChar {
			totalScore += score
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	stackSize := len(stack)
	str = make([]rune, stackSize)
	for i := stackSize - 1; i >= 0; i-- {
		str[i] = stack[len(stack)-1-i]
	}

	return totalScore
}
