package string

import (
	"fmt"
	"testing"
)

func TestAddStrings(t *testing.T) {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
	fmt.Println(addStrings("0", "0"))
}
