package tree

type TreeSum struct {
	isSumTree bool
	sum       int
}

func isSumTree(root *TreeNode) *TreeSum {

	// Base case
	if root == nil {
		return &TreeSum{true, 0}
	}
	if root.Left == nil && root.Right == nil {
		return &TreeSum{true, root.Data}
	}

	left := isSumTree(root.Left)
	right := isSumTree(root.Right)

	isLeftSumTree := left.isSumTree
	isRightSumTree := right.isSumTree
	isSumEqual := root.Data == (left.sum + right.sum)

	ans := &TreeSum{false, 0}

	if isLeftSumTree && isRightSumTree && isSumEqual {
		ans.isSumTree = true
		ans.sum = root.Data + left.sum + right.sum // 2 * root.data
	}

	return ans
}
