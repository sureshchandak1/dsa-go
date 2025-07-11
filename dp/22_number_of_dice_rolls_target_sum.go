package dp

import "slices"

func noOfWays(face, dice, target int64) int64 {
	// return noOfWaysRec(face, dice, target)

	// dp := make([][]int64, dice+1)
	// for i := range dp {
	// 	dp[i] = slices.Repeat([]int64{-1}, int(target)+1)
	// }
	// return noOfWaysMem(face, dice, target, dp)

	// return noOfWaysTab(face, dice, target)

	return noOfWaysSC(face, dice, target)
}

// Recursion
func noOfWaysRec(face, dice, target int64) int64 {

	// Base case
	if target < 0 {
		return 0
	}
	if dice != 0 && target == 0 {
		return 0
	}
	if dice == 0 && target != 0 {
		return 0
	}
	if dice == 0 && target == 0 {
		return 1
	}

	var ans int64
	for i := int64(1); i <= face; i++ {
		ans = (ans%int64(MOD) + noOfWaysRec(face, dice-1, target-i)%int64(MOD)) % int64(MOD)
	}

	return ans
}

// Recursion + Memorization
// T.C = O(d * f * t), S.C = O(d * t)
func noOfWaysMem(face, dice, target int64, dp [][]int64) int64 {

	// Base case
	if target < 0 {
		return 0
	}
	if dice != 0 && target == 0 {
		return 0
	}
	if dice == 0 && target != 0 {
		return 0
	}
	if dice == 0 && target == 0 {
		return 1
	}

	if dp[dice][target] != int64(-1) {
		return dp[dice][target]
	}

	var ans int64
	for i := int64(1); i <= face; i++ {
		ans = (ans%int64(MOD) + noOfWaysMem(face, dice-1, target-i, dp)%int64(MOD)) % int64(MOD)
	}

	dp[dice][target] = ans

	return dp[dice][target]
}

// Tabulation
// T.C = O(d * f * t), S.C = O(d * t)
func noOfWaysTab(f, d, t int64) int64 {

	dp := make([][]int64, d+1)
	for i := range dp {
		dp[i] = slices.Repeat([]int64{0}, int(t)+1)
	}

	// Base case analysing
	dp[0][0] = 1

	for dice := int64(1); dice <= d; dice++ {
		for target := int64(1); target <= t; target++ {

			var ans int64
			for i := int64(1); i <= f; i++ {
				if target-i >= 0 {
					ans = (ans%int64(MOD) + dp[dice-1][target-i]%int64(MOD)) % int64(MOD)
				}
			}

			dp[dice][target] = ans

		}
	}

	return dp[d][t]

}

// Space Optimization
// T.C = O(d * f * t), S.C = O(t)
func noOfWaysSC(f, d, t int64) int64 {

	prev := make([]int64, t+1)
	curr := make([]int64, t+1)

	prev[0] = 1

	for dice := int64(1); dice <= d; dice++ {
		for target := int64(1); target <= t; target++ {

			var ans int64
			for i := int64(1); i <= f; i++ {
				if target-i >= 0 {
					ans = (ans%int64(MOD) + prev[target-i]%int64(MOD)) % int64(MOD)
				}
			}

			curr[target] = ans

		}

		prev = slices.Clone(curr)
	}

	return prev[t]
}

/**
 *    ( a + b) % c = ( ( a % c ) + ( b % c ) ) % c
 *    ( a * b) % c = ( ( a % c ) * ( b % c ) ) % c
 *    ( a – b) % c = ( ( a % c ) – ( b % c ) ) % c
 */
