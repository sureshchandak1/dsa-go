package tree

import "math"

type TreeBalanced struct {
	isBalanced bool
	height     int
}

func isBalancedTree(root *TreeNode) *TreeBalanced {

	// Base case
	if root == nil {
		return &TreeBalanced{true, 0}
	}

	left := isBalancedTree(root.Left)
	right := isBalancedTree(root.Right)

	leftAns := left.isBalanced
	rightAns := right.isBalanced
	diff := math.Abs(float64(left.height)-float64(right.height)) <= 1

	ans := &TreeBalanced{}
	ans.height = max(left.height, right.height) + 1

	if leftAns && rightAns && diff {
		ans.isBalanced = true
	} else {
		ans.isBalanced = false
	}

	return ans

}
