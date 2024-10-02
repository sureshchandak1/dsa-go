package recursion

import "fmt"

func climbStairWays(nStairs int) int {

	// Base case
	if nStairs < 0 {
		return 0
	}
	if nStairs == 0 {
		return 1
	}

	count1Stair := climbStairWays(nStairs - 1)
	count2Stair := climbStairWays(nStairs - 2)

	return count1Stair + count2Stair
}

func reachHome(startPosition int, destination int) {

	fmt.Printf("Position: %d, Destination: %d\n", startPosition, destination)

	// Base case
	if startPosition == destination {
		fmt.Println("Reach Home")
		return
	}

	// Recursion
	reachHome(startPosition+1, destination)
}
