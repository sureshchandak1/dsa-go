package search

import (
	"fmt"
	"testing"
)

func TestBookAllocation(t *testing.T) {

	bookPages := []int{10, 20, 30, 40}
	totalStudent := 2
	fmt.Println(bookAllocation(bookPages, totalStudent))

	bookPages = []int{12, 34, 67, 90}
	totalStudent = 2
	fmt.Println(bookAllocation(bookPages, totalStudent))

	bookPages = []int{25, 46, 28, 49, 24}
	totalStudent = 4
	fmt.Println(bookAllocation(bookPages, totalStudent))

	bookPages = []int{2, 2, 3, 3, 4, 4, 1}
	totalStudent = 4
	fmt.Println(bookAllocation(bookPages, totalStudent))

}
