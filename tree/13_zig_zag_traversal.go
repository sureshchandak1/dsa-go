package tree

func zigZagTraversal(root *TreeNode) []int {

	result := []int{}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	leftToRight := true

	for len(queue) != 0 {

		size := len(queue)

		ans := make([]int, size)

		for i := 0; i < size; i++ {

			front := queue[0]
			queue = queue[1:] // remove front node

			index := 0
			if leftToRight {
				index = i
			} else {
				index = size - i - 1
			}

			ans[index] = front.Data

			if front.Left != nil {
				queue = append(queue, front.Left)
			}

			if front.Right != nil {
				queue = append(queue, front.Right)
			}

		}

		leftToRight = !leftToRight

		result = append(result, ans...)

	}

	return result
}
