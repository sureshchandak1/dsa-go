package tree

import "fmt"

/*
*      N L R
* */
func preOrderTraversal(root *TreeNode) {

	// Base case
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Data) // N = Data

	preOrderTraversal(root.Left) // L = Left

	preOrderTraversal(root.Right) // R = Right

}
