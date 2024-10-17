package queue

import (
	"fmt"
	"testing"
)

func TestFirstNonRepeating(t *testing.T) {
	fmt.Println(firstNonRepeating("aabc"))
	fmt.Println(firstNonRepeating("zz"))
	fmt.Println(firstNonRepeating("ababc"))
}
