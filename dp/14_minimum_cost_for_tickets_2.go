package dp

// type Pair[T, U any] struct {
// 	First  T
// 	Second U
// }

// type Pair struct {
// 	a, b any
// }

type ticketPair struct {
	day, cost int
}

// T.C = O(n), S.C = O(1)
func minCostTicketsOptimized(days, cost []int) int {

	week := make([]ticketPair, 0)
	month := make([]ticketPair, 0)

	ans := 0

	for _, day := range days {

		// remove expired days
		for len(week) > 0 && week[0].day+7 <= day {
			week = week[1:] // remove top
		}
		for len(month) > 0 && month[0].day+30 <= day {
			month = month[1:] // remove top
		}

		week = append(week, ticketPair{day, ans + cost[1]})
		month = append(month, ticketPair{day, ans + cost[2]})

		ans = min(ans+cost[0], min(week[0].cost, month[0].cost))
	}

	return ans
}
