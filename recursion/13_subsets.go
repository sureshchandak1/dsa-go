package recursion

func getSubsets(nums []int) [][]int {
	result := [][]int{}
	subsets(nums, []int{}, 0, &result)
	return result
}

func subsets(nums []int, output []int, index int, result *[][]int) {

	// Base case
	if index >= len(nums) {
		*result = append(*result, output)
		return
	}

	// exclude
	subsets(nums, output, index+1, result)

	// include
	output = append(output, nums[index])
	subsets(nums, output, index+1, result)
}
