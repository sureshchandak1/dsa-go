package tree

import "fmt"

func levelOrderTraversal(root *TreeNode) []int {

	result := []int{}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front node

		result = append(result, front.Data)

		if front.Left != nil {
			queue = append(queue, front.Left)
		}

		if front.Right != nil {
			queue = append(queue, front.Right)
		}

	}

	return result
}

func levelOrderTraversal2(root *TreeNode) {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	queue = append(queue, nil) // null - separator

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front node

		if front == nil {
			fmt.Println()
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			fmt.Printf("%d ", front.Data)

			if front.Left != nil {
				queue = append(queue, front.Left)
			}

			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}

	}

}
