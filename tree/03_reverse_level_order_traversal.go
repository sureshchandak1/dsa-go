package tree

func reverseLevelOrderTraversal(root *TreeNode) []int {

	result := []int{}

	stack := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front node

		stack = append(stack, front.Data)

		if front.Left != nil {
			queue = append(queue, front.Left)
		}

		if front.Right != nil {
			queue = append(queue, front.Right)
		}

	}

	for len(stack) != 0 {
		size := len(stack)
		top := stack[size-1]   // get top
		stack = stack[:size-1] // remove top

		result = append(result, top)
	}

	return result
}
