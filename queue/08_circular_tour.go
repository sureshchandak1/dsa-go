package queue

func canCompleteCircuit(gas []int, cost []int) int {

	deficit := 0
	balance := 0
	start := 0

	for i := 0; i < len(gas); i++ {

		balance = balance + gas[i] - cost[i]

		if balance < 0 {
			deficit = deficit + balance
			start = i + 1
			balance = 0
		}
	}

	if balance+deficit >= 0 {
		return start
	}

	return -1

}
