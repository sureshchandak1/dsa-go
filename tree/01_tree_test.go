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

func TestBoundaryTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(boundaryTraversal(root))
}

func TestVerticalTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(verticalTraversal(root))
}

func TestBinaryTreeTopView(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(binaryTreeTopView(root))
}

func TestBinaryTreeBottomView(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(binaryTreeBottomView(root))
}

func TestBinaryTreeLeftView(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(binaryTreeLeftView(root))
}

func TestBinaryTreeRightView(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(binaryTreeRightView(root))
}

func TestSumOfLongRootToLeafPath(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(sumOfLongRootToLeafPath(root))
}

func TestSumRootToLeaf(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(sumRootToLeaf(root))
}

func TestLowestCommonAncestor(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	result := lowestCommonAncestor(root, &TreeNode{Data: 40}, &TreeNode{Data: 50})
	if result != nil {
		fmt.Println(result.Data)
	}
}

func TestPathSum(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(pathSum(root, 140))
}

func TestGetKthAncestor(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(getKthAncestor(root, 2, 50))
}

func TestMaximumSumOfNodes(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	fmt.Println(maximumSumOfNodes(root))
}

func TestBuildBinaryTreeFromPreorderInorder(t *testing.T) {
	root := buildBinaryTreeFromPreorderInorder([]int{3, 1, 4, 0, 5, 2}, []int{0, 1, 3, 4, 2, 5})
	postOrderTraversal(root)
}

func TestBuildBinaryTreeFromPostorderInorder(t *testing.T) {
	root := buildBinaryTreeFromPostorderInorder([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	preOrderTraversal(root)
}

func TestBurnTree(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	result := minTimeToBurnTree(root, 40)
	fmt.Println(result)
}

func TestMorrisTraversal(t *testing.T) {
	root := buildTreeString("10 20 30 40 60 N N N N 50 N N")
	result := morrisTraversal(root)
	fmt.Println(result)
	inorderTraversal(root)
}
