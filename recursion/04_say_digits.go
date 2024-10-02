package recursion

import "fmt"

func sayDigits(number int, counts []string) {

	// Base case
	if number == 0 {
		return
	}

	// Processing
	digit := number % 10

	// Recursion
	sayDigits(number/10, counts)

	fmt.Print(" ", counts[digit])

}
