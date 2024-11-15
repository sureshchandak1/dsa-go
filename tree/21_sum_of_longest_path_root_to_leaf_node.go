package tree

/*
   https://www.geeksforgeeks.org/problems/sum-of-the-longest-bloodline-of-a-tree/1
*/
func sumOfLongRootToLeafPath(root *TreeNode) int {

	var maxLen int = 0
	var maxSum int = 0

	rootToLeafSum(root, 0, 0, &maxSum, &maxLen)

	return maxSum
}

func rootToLeafSum(node *TreeNode, sum, len int, maxSum *int, maxLen *int) {

	// Base case
	if node == nil {
		if len > *maxLen {
			*maxLen = len
			*maxSum = sum
		} else if len == *maxLen {
			*maxSum = max(sum, *maxSum)
		}

		return
	}

	sum = sum + node.Data

	rootToLeafSum(node.Left, sum, len+1, maxSum, maxLen)

	rootToLeafSum(node.Right, sum, len+1, maxSum, maxLen)

}
