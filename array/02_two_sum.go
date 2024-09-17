package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

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
