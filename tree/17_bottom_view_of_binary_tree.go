package tree

type BottomViewPair struct {
	node               *TreeNode
	horizontalDistance int
}

func binaryTreeBottomView(root *TreeNode) []int {

	result := []int{}

	if root == nil {
		return result
	}

	topNode := make(map[int]int)

	queue := []*BottomViewPair{}
	queue = append(queue, &BottomViewPair{root, 0})

	for len(queue) != 0 {

		front := queue[0]
		queue = queue[1:] // remove front

		node := front.node
		hd := front.horizontalDistance

		topNode[hd] = node.Data

		if node.Left != nil {
			queue = append(queue, &BottomViewPair{node.Left, hd - 1})
		}

		if node.Right != nil {
			queue = append(queue, &BottomViewPair{node.Right, hd + 1})
		}
	}

	for _, value := range topNode {
		result = append(result, value)
	}

	return result

}
