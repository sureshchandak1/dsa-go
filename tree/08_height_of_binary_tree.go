package tree

func treeHeight(root *TreeNode) int {

	// Base case
	if root == nil {
		return 0
	}

	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)

	return max(leftHeight, rightHeight) + 1

}
