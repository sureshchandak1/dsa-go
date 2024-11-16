package tree

import "math"

func getKthAncestor(root *TreeNode, k int, targetNodeVal int) int {

	result := findKthAncestor(root, &k, &targetNodeVal)

	if result == nil || result.Data == targetNodeVal {
		return -1
	}

	return result.Data
}

func findKthAncestor(root *TreeNode, k *int, node *int) *TreeNode {

	// Base case
	if root == nil {
		return nil
	}

	if root.Data == *node {
		return root
	}

	leftAns := findKthAncestor(root.Left, k, node)

	rightAns := findKthAncestor(root.Right, k, node)

	if leftAns != nil && rightAns == nil {
		*k--
		if *k <= 0 {
			*k = int(math.Pow(2, 30))
			return root
		}
		return leftAns
	}

	if leftAns == nil && rightAns != nil {
		*k--
		if *k <= 0 {
			*k = int(math.Pow(2, 30))
			return root
		}
		return rightAns
	}

	return nil
}
