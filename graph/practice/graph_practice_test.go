package practice

import (
	"fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	fmt.Println(findJudge(2, [][]int{{1, 2}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	fmt.Println(findJudgeOptimized(2, [][]int{{1, 2}}))
	fmt.Println(findJudgeOptimized(3, [][]int{{1, 3}, {2, 3}}))
	fmt.Println(findJudgeOptimized(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
