package leetcode

func closestPrimes(left int, right int) []int {

	prime := sieveAlgo(right)

	result := []int{-1, -1}
	minDiff := 2147483647
	prev := -1

	for i := left; i <= right; i++ {
		if prime[i] {
			if prev == -1 {
				prev = i
			} else {
				if i-prev < minDiff {
					result[0] = prev
					result[1] = i
					minDiff = i - prev
				}

				prev = i
			}
		}
	}

	return result
}

// Sieve Algorithm for find prime number in range
func sieveAlgo(right int) []bool {

	prime := make([]bool, right+1)
	for i := range prime {
		prime[i] = true
	}

	prime[0] = false
	prime[1] = false

	for p := 2; p*p <= right; p++ {
		if prime[p] {
			for i := p * p; i <= right; i += p {
				prime[i] = false
			}
		}
	}

	return prime
}
