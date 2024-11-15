package tree

/**

https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// Base case
	if root == nil {
		return nil
	}
	if root.Data == p.Data || root.Data == q.Data {
		return root
	}

	leftAns := lowestCommonAncestor(root.Left, p, q)

	rightAns := lowestCommonAncestor(root.Right, p, q)

	if leftAns != nil && rightAns != nil {
		return root
	} else if leftAns != nil && rightAns == nil {
		return leftAns
	} else if leftAns == nil && rightAns != nil {
		return rightAns
	} else {
		return nil
	}

}
