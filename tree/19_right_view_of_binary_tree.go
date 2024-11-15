package tree

func binaryTreeRightView(root *TreeNode) []int {

	result := []int{}

	rightView(root, 0, &result)

	return result
}

func rightView(node *TreeNode, level int, result *[]int) {

	// Base case
	if node == nil {
		return
	}

	// entered into new level
	if level == len(*result) {
		*result = append(*result, node.Data)
	}

	rightView(node.Right, level+1, result)

	rightView(node.Left, level+1, result)

}
