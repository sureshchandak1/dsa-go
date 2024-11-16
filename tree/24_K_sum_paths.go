package tree

/**

  https://leetcode.com/problems/path-sum-iii/description/

*/

func pathSum(root *TreeNode, targetSum int) int {

	count := 0
	path := []int{}

	solve(root, targetSum, &count, &path)

	return count
}

func solve(node *TreeNode, targetSum int, count *int, path *[]int) {

	// Base case
	if node == nil {
		return
	}

	*path = append(*path, node.Data)

	solve(node.Left, targetSum, count, path)

	solve(node.Right, targetSum, count, path)

	sum := 0
	for i := len(*path) - 1; i >= 0; i-- {
		sum = sum + (*path)[i]
		if sum == targetSum {
			*count++
		}
	}

	// Remove last added element
	*path = (*path)[:len(*path)-1]
}
