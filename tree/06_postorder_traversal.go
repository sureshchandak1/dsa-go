package tree

import "fmt"

/*
*      L R N
* */
func postOrderTraversal(root *TreeNode) {

	// Base case
	if root == nil {
		return
	}

	postOrderTraversal(root.Left) // L = Left

	postOrderTraversal(root.Right) // R = Right

	fmt.Printf("%d ", root.Data)

}
