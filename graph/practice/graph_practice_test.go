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

func TestWordExist(t *testing.T) {
	fmt.Println(wordExist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(wordExist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(wordExist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
}

func TestIslandPerimeter(t *testing.T) {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1}}))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
}

func TestNumIslands(t *testing.T) {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}

func TestFindFarmland(t *testing.T) {
	fmt.Println(findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Println(findFarmland([][]int{{1, 1}, {1, 1}}))
	fmt.Println(findFarmland([][]int{{0}}))
}

func TestValidPath(t *testing.T) {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}

func TestFindRotateSteps(t *testing.T) {
	fmt.Println(findRotateSteps("godding", "gd"))
	fmt.Println(findRotateSteps("godding", "gg"))
	fmt.Println(findRotateSteps("godding", "godding"))
}

func TestMinOperationsForXORtoK(t *testing.T) {
	fmt.Println(minOperationsForXORtoK([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperationsForXORtoK([]int{2, 0, 2, 0}, 0))
}

func TestMaximumGold(t *testing.T) {
	fmt.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
	fmt.Println(getMaximumGold([][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}))
}
