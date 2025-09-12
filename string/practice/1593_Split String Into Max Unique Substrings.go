package practice

func maxUniqueSplit(s string) int {

	maxCount := 0
	set := make(map[string]struct{})

	solveMaxUniqueSplit(s, set, 0, &maxCount)

	return maxCount
}

func solveMaxUniqueSplit(s string, set map[string]struct{}, index int, maxCount *int) {

	n := len(s)

	// Base case
	if index == n {
		*maxCount = max(*maxCount, len(set))
		return
	}

	for i := index; i < n; i++ {
		// Check
		sub := s[index : i+1]
		if _, exists := set[sub]; !exists {
			set[sub] = struct{}{}

			solveMaxUniqueSplit(s, set, i+1, maxCount)

			delete(set, sub)
		}
	}
}
