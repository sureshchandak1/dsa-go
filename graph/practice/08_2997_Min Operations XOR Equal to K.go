package practice

func minOperationsForXORtoK(nums []int, k int) int {

	// Find XOR of array
	finalXOR := 0
	for _, num := range nums {
		finalXOR ^= num
	}

	// find the different bits
	count := 0
	for k > 0 || finalXOR > 0 {

		if k%2 != finalXOR%2 {
			// check last bit
			count++
		}

		// Remove last bit
		k /= 2
		finalXOR /= 2
	}

	return count
}
