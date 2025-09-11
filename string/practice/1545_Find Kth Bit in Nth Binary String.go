package practice

import "math"

func findKthBit(n int, k int) string {
	len := int(math.Pow(2, float64(n))) - 1
	return string(solveFindKthBit(len, k))
}

func solveFindKthBit(len int, k int) byte {

	// Base case
	if len == 1 {
		return '0'
	}

	half := len / 2
	middle := half + 1

	if k == middle {
		return '1'
	} else if k < middle { // left half
		return solveFindKthBit(half, k)
	} else {
		// right half
		ans := solveFindKthBit(half, len-k+1)
		if ans == '0' {
			return '1'
		}
		return '0'
	}
}
