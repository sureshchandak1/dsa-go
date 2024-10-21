package tree

import "fmt"

/*
*      L N R
* */
func inorderTraversal(root *TreeNode) {

	// Base case
	if root == nil {
		return
	}

	inorderTraversal(root.Left) // L = Left

	fmt.Printf("%d ", root.Data) // N = Data

	inorderTraversal(root.Right) // R = Right

}
