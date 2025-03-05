package array

func twoSum(nums []int, target int) []int {

	indexMap := make(map[int]int)

	for index, num := range nums {

		diff := target - num

		value, exists := indexMap[diff]
		if exists {
			return []int{value, index}
		}

		indexMap[num] = index
	}

	return []int{}

}
