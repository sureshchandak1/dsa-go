package tree

func boundaryTraversal(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	result = append(result, root.Data)

	// Tree left node store in result
	traverseLeft(root.Left, &result)

	// Tree leaf node store in result
	traverseLeaf(root.Left, &result)  // left leaf node
	traverseLeaf(root.Right, &result) // right leaf node

	// Tree right node store in result
	traverseRight(root.Right, &result)

	return result
}

func traverseLeft(root *TreeNode, result *[]int) {

	// Base case
	if root == nil {
		return
	}
	// ignore leaf node
	if root.Left == nil && root.Right == nil {
		return
	}

	*result = append(*result, root.Data)

	if root.Left != nil {
		traverseLeft(root.Left, result)
	} else {
		traverseLeft(root.Right, result)
	}
}

func traverseLeaf(root *TreeNode, result *[]int) {

	// Base case
	if root == nil {
		return
	}

	// Leaf node
	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Data)
		return
	}

	traverseLeaf(root.Left, result)

	traverseLeaf(root.Right, result)

}

func traverseRight(root *TreeNode, result *[]int) {

	// Base case
	if root == nil {
		return
	}
	// ignore leaf node
	if root.Left == nil && root.Right == nil {
		return
	}

	if root.Right != nil {
		traverseRight(root.Right, result)
	} else {
		traverseRight(root.Left, result)
	}

	*result = append(*result, root.Data)

}
