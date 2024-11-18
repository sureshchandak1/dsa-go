package tree

func buildBinaryTreeFromPostorderInorder(inorder []int, postorder []int) *TreeNode {

	mapping := make(map[int]int)
	for index, value := range inorder {
		mapping[value] = index
	}

	postorderIndex := len(postorder) - 1
	root := createTreePostIn(inorder, postorder, &postorderIndex, 0, len(inorder)-1, mapping)
	return root
}

func createTreePostIn(inorder []int, postorder []int, postorderIndex *int, inorderStart, inorderEnd int, mapping map[int]int) *TreeNode {

	// Base case
	if *postorderIndex < 0 || inorderStart > inorderEnd {
		return nil
	}

	element := postorder[*postorderIndex]
	*postorderIndex--
	root := &TreeNode{Data: element}

	//position := findPosition(inorder, element)
	position := mapping[element]

	root.Right = createTreePostIn(inorder, postorder, postorderIndex, position+1, inorderEnd, mapping)

	root.Left = createTreePostIn(inorder, postorder, postorderIndex, inorderStart, position-1, mapping)

	return root
}
