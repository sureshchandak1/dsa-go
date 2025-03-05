package tree

/**

https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

*/

func buildBinaryTreeFromPreorderInorder(inorder []int, preorder []int) *TreeNode {

	mapping := make(map[int]int)
	for index, value := range inorder {
		mapping[value] = index
	}

	preorderIndex := 0
	root := createTree(inorder, preorder, &preorderIndex, 0, len(inorder)-1, mapping)
	return root
}

func createTree(inorder []int, preorder []int, preorderIndex *int, inorderStart, inorderEnd int, mapping map[int]int) *TreeNode {

	// Base case
	if *preorderIndex >= len(preorder) || inorderStart > inorderEnd {
		return nil
	}

	element := preorder[*preorderIndex]
	*preorderIndex++
	root := &TreeNode{Data: element}

	//position := findPosition(inorder, element)
	position := mapping[element]

	root.Left = createTree(inorder, preorder, preorderIndex, inorderStart, position-1, mapping)

	root.Right = createTree(inorder, preorder, preorderIndex, position+1, inorderEnd, mapping)

	return root
}

/*
func findPosition(inorder []int, element int) int {

	for index, value := range inorder {
		if value == element {
			return index
		}
	}

	return -1
}
*/
