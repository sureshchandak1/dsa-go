package tree

func countLeafNodes(root *TreeNode) int {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	count := 0

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front node

		if front.Left != nil {
			queue = append(queue, front.Left)
		}

		if front.Right != nil {
			queue = append(queue, front.Right)
		}

		if front.Left == nil && front.Right == nil {
			count++
		}
	}

	return count

}

// L N R
func inorderCountLeafNodes(root *TreeNode, count int) int {

	// Base case
	if root == nil {
		return count
	}

	// L
	leftLeafNodeCount := inorderCountLeafNodes(root.Left, count)

	// N
	if root.Left == nil && root.Right == nil {
		count++
	}

	// R
	rightLeafNodeCount := inorderCountLeafNodes(root.Right, count)

	return leftLeafNodeCount + rightLeafNodeCount

}
