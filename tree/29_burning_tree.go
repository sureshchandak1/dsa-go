package tree

func minTimeToBurnTree(root *TreeNode, target int) int {

	nodeToParent := make(map[*TreeNode]*TreeNode)

	targetNode := createParentMapping(root, target, nodeToParent)

	ans := burnTree(targetNode, nodeToParent)

	return ans
}

func burnTree(root *TreeNode, nodeToParent map[*TreeNode]*TreeNode) int {

	visited := make(map[*TreeNode]bool)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	visited[root] = true

	var ans = 0

	for len(queue) != 0 {

		flag := false // node add in queue or not
		size := len(queue)

		for i := 0; i < size; i++ {

			front := queue[0]
			queue = queue[1:] // remove front node

			if front.Left != nil {
				_, exits := visited[front.Left]
				if !exits {
					flag = true
					visited[front.Left] = true
					queue = append(queue, front.Left)
				}
			}

			if front.Right != nil {
				_, exits := visited[front.Right]
				if !exits {
					flag = true
					visited[front.Right] = true
					queue = append(queue, front.Right)
				}
			}

			if nodeToParent[front] != nil {
				_, exits := visited[nodeToParent[front]]
				if !exits {
					flag = true
					visited[nodeToParent[front]] = true
					queue = append(queue, nodeToParent[front])
				}
			}
		}

		if flag {
			ans++
		}

	}

	return ans
}

func createParentMapping(root *TreeNode, target int, nodeToParent map[*TreeNode]*TreeNode) *TreeNode {

	var targetNode *TreeNode

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	nodeToParent[root] = nil

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front node

		if target == front.Data {
			targetNode = front
		}

		if front.Left != nil {
			nodeToParent[front.Left] = front
			queue = append(queue, front.Left)
		}

		if front.Right != nil {
			nodeToParent[front.Right] = front
			queue = append(queue, front.Right)
		}
	}

	return targetNode
}
