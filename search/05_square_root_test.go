package search

import (
	"fmt"
	"testing"
)

func TestSquareRoot(t *testing.T) {

	fmt.Println(morePrecision(36, 3, float64(squareRoot(36))))
	fmt.Println(squareRoot(28))
	fmt.Println(morePrecision(101, 3, float64(squareRoot(101))))
	fmt.Println(squareRoot(10000000000))

}
