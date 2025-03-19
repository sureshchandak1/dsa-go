package slidingwindow

import (
	"fmt"
	"testing"
)

func TestMinOperationsToEqual1(t *testing.T) {
	fmt.Println(minOperationsToEqual1([]int{0, 1, 1, 1, 0, 0}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 1, 1}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 0, 1, 0}))
	fmt.Println(minOperationsToEqual1([]int{0, 1, 0, 0, 1, 0, 1}))
	fmt.Println(minOperationsToEqual1([]int{1, 1, 0, 1}))
}
