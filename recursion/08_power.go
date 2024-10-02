package recursion

func powerOptimized(number, power int) int {

	// Base case
	if power == 0 {
		return 1
	}
	if power == 1 {
		return number
	}

	halfResult := powerOptimized(number, power/2) // half power result

	if power%2 == 0 {
		return halfResult * halfResult
	} else {
		return number * halfResult * halfResult
	}

}
