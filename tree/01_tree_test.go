package tree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	buildTreeString("10 20 30 40 60 N N N N 50 N N N")
}

func TestLevelOrderTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(levelOrderTraversal(root))
	levelOrderTraversal2(root)
}

func TestReverseLevelOrderTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(reverseLevelOrderTraversal(root))
}

func TestInPrePostOrderTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	inorderTraversal(root)
	fmt.Println()
	preOrderTraversal(root)
	fmt.Println()
	postOrderTraversal(root)
	fmt.Println()
}

func TestCountLeafNodes(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(countLeafNodes(root))
	fmt.Println(inorderCountLeafNodes(root, 0))
}

func TestTreeHeight(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(treeHeight(root))
}

func TestTreeDiameter(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(treeDiameter(root).diameter)
}

func TestTreeBalanced(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(isBalancedTree(root).isBalanced)
}

func TestIdenticalTrees(t *testing.T) {
	root1 := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	root2 := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(identicalTrees(root1, root2))
}

func TestIsSumTree(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(isSumTree(root).isSumTree)
}

func TestZigZagTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(zigZagTraversal(root))
}
