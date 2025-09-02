package practice

import "strconv"

func countSeniorsOptimized(details []string) int {
	count := 0

	for _, detail := range details {

		a := int(detail[11] - '0')
		b := int(detail[12] - '0')
		ageVal := a*10 + b

		if ageVal > 60 {
			count++
		}
	}

	return count
}

func countSeniors(details []string) int {
	count := 0

	for _, detail := range details {

		age := detail[11:13]
		ageVal, _ := strconv.Atoi(age)

		if ageVal > 60 {
			count++
		}
	}

	return count
}
