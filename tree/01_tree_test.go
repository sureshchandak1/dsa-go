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
