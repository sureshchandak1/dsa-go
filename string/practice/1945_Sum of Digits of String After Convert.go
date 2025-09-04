package practice

func getSumStringDigits(s string, k int) int {

	sum := 0
	result := 0

	for _, ch := range s {
		num := int(ch - 'a' + 1)

		for num > 0 {
			sum += num % 10
			num /= 10
		}
	}

	result = sum

	for k > 1 {

		digitSum := 0
		for result > 0 {
			digitSum += result % 10
			result /= 10
		}

		k--
		result = digitSum
	}

	return result

}
