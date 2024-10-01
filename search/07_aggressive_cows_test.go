package search

import (
	"fmt"
	"testing"
)

func TestAggressiveCows(t *testing.T) {

	stalls := []int{1, 2, 4, 8, 9}
	totalCows := 3
	fmt.Println(aggressiveCows(stalls, totalCows))

	stalls = []int{10, 1, 2, 7, 5}
	totalCows = 3
	fmt.Println(aggressiveCows(stalls, totalCows))

	stalls = []int{0, 3, 4, 7, 10, 9}
	totalCows = 4
	fmt.Println(aggressiveCows(stalls, totalCows))

	stalls = []int{4, 2, 1, 3, 6}
	totalCows = 2
	fmt.Println(aggressiveCows(stalls, totalCows))
}
