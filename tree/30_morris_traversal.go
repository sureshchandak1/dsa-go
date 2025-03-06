package tree

func morrisTraversal(root *TreeNode) []int {

	result := []int{}

	var pre *TreeNode
	current := root
	for current != nil {

		if current.Left == nil {
			result = append(result, current.Data)
			current = current.Right
		} else {
			// Find the inorder predecessor of current
			pre = current.Left
			for pre.Right != nil && pre.Right != current {
				pre = pre.Right
			}

			/* Make current as right
			   child of its
			 * inorder predecessor */
			if pre.Right == nil {
				pre.Right = current
				current = current.Left
			} else {
				pre.Right = nil
				result = append(result, current.Data)
				current = current.Right
			}
		}
	}

	return result
}
