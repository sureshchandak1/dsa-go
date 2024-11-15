package tree

func binaryTreeLeftView(root *TreeNode) []int {

	result := []int{}

	leftView(root, 0, &result)

	return result
}

func leftView(node *TreeNode, level int, result *[]int) {

	// Base case
	if node == nil {
		return
	}

	// entered into new level
	if level == len(*result) {
		*result = append(*result, node.Data)
	}

	leftView(node.Left, level+1, result)

	leftView(node.Right, level+1, result)
}
