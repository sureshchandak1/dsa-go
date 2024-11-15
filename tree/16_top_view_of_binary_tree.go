package tree

type TopViewPair struct {
	node               *TreeNode
	horizontalDistance int
}

func binaryTreeTopView(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	topNode := make(map[int]int)

	queue := []*TopViewPair{}
	queue = append(queue, &TopViewPair{root, 0})

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front

		node := front.node
		hd := front.horizontalDistance

		_, exists := topNode[hd]
		if !exists {
			topNode[hd] = node.Data
		}

		if node.Left != nil {
			queue = append(queue, &TopViewPair{node.Left, hd - 1})
		}

		if node.Right != nil {
			queue = append(queue, &TopViewPair{node.Right, hd + 1})
		}
	}

	for _, value := range topNode {
		result = append(result, value)
	}

	return result
}
