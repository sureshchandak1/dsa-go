package search

import (
	"fmt"
	"testing"
)

func TestPainterPartition(t *testing.T) {

	boards := []int{5, 5, 5, 5}
	totalPainter := 2
	fmt.Println(painterPartition(boards, totalPainter))

	boards = []int{2, 1, 5, 6, 2, 3}
	totalPainter = 2
	fmt.Println(painterPartition(boards, totalPainter))

	boards = []int{10, 20, 30, 40}
	totalPainter = 2
	fmt.Println(painterPartition(boards, totalPainter))

	boards = []int{48, 90}
	totalPainter = 2
	fmt.Println(painterPartition(boards, totalPainter))

	boards = []int{5, 10, 30, 20, 15}
	totalPainter = 3
	fmt.Println(painterPartition(boards, totalPainter))
}
