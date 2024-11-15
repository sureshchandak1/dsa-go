package tree

/**

https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/description/

*/

func sumRootToLeaf(root *TreeNode) int {
	return depthFirstSearch(root, 0)
}

func depthFirstSearch(node *TreeNode, sum int) int {

	// Base case
	if node == nil {
		return 0
	}

	sum = (sum << 1) | node.Data

	if node.Left == nil && node.Right == nil {
		return sum
	}

	leftSum := depthFirstSearch(node.Left, sum)

	rightSum := depthFirstSearch(node.Right, sum)

	return leftSum + rightSum
}
