package tree

func flattenTree(root *TreeNode) {

	curr := root

	for curr != nil {

		if curr.Left != nil {

			pred := curr.Left
			for pred != nil && pred.Right != nil {
				pred = pred.Right
			}

			pred.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}

		curr = curr.Right
	}
}
