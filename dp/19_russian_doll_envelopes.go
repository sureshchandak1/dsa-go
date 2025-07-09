package dp

import "sort"

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[j][1] < envelopes[i][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// heights := make([]int, len(envelopes))
	// for i := range envelopes {
	// 	heights[i] = envelopes[i][1]
	// }

	return maxEnvelopesBinarySearch(envelopes)
}

func maxEnvelopesBinarySearch(envelopes [][]int) int {

	n := len(envelopes)

	if n == 0 {
		return 0
	}

	ans := []int{}
	ans = append(ans, envelopes[0][1])

	for i := 1; i < n; i++ {

		if envelopes[i][1] > ans[len(ans)-1] {
			ans = append(ans, envelopes[i][1])
		} else {
			index := lowerBound(ans, envelopes[i][1])
			ans[index] = envelopes[i][1]
		}
	}

	return len(ans)

}
