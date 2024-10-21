package tree

type TreeDiameter struct {
	diameter int
	height   int
}

func treeDiameter(root *TreeNode) *TreeDiameter {

	// Base case
	if root == nil {
		return &TreeDiameter{0, 0}
	}

	left := treeDiameter(root.Left)
	right := treeDiameter(root.Right)

	op1 := left.diameter
	op2 := right.diameter
	op3 := left.height + right.height

	ans := &TreeDiameter{}
	ans.diameter = max(op1, max(op2, op3))
	ans.height = max(left.height, right.height) + 1

	return ans
}
