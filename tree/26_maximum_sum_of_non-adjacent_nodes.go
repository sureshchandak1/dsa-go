package tree

/**

https://www.geeksforgeeks.org/problems/maximum-sum-of-non-adjacent-nodes/1

*/

type PairMaxSum struct {
	includeSum int
	excludeSum int
}

func maximumSumOfNodes(root *TreeNode) int {

	result := getMaxSum(root)

	return max(result.includeSum, result.excludeSum)
}

func getMaxSum(root *TreeNode) PairMaxSum {

	// Base case
	if root == nil {
		return PairMaxSum{0, 0}
	}

	leftAns := getMaxSum(root.Left)

	rightAns := getMaxSum(root.Right)

	result := PairMaxSum{0, 0}

	result.includeSum = root.Data + leftAns.excludeSum + rightAns.excludeSum

	result.excludeSum = max(leftAns.includeSum, leftAns.excludeSum) + max(rightAns.includeSum, rightAns.excludeSum)

	return result

}
