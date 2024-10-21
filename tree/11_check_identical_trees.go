package tree

func identicalTrees(root1 *TreeNode, root2 *TreeNode) bool {

	// Base case
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1 != nil && root2 == nil {
		return false
	}

	left := identicalTrees(root1.Left, root2.Left)
	right := identicalTrees(root1.Right, root2.Right)
	isSameValue := root1.Data == root2.Data

	return left && right && isSameValue

}
