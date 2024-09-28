package search

func squareRoot(number int) int {

	start := 0
	end := number
	mid := start + (end-start)/2

	ans := -1

	for start <= end {

		squ := mid * mid
		if number == squ {
			return mid
		}

		if squ > number {
			end = mid - 1
		} else {
			ans = mid
			start = mid + 1
		}

		mid = start + (end-start)/2
	}

	return ans
}

func morePrecision(number float64, precision int, tempSol float64) float64 {

	var factor float64 = 1
	var ans float64 = tempSol

	for i := 0; i < precision; i++ {

		factor = factor / 10

		for j := ans; j*j < number; j = j + factor {
			ans = j
		}
	}

	return ans
}
