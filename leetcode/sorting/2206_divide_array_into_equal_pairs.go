package sorting

import "sort"

func divideArray(nums []int) bool {
	return divideArraySolution2(nums)
}

func divideArraySolution1(nums []int) bool {
	set := make(map[int]bool)

	for _, num := range nums {
		_, ok := set[num]

		if ok {
			delete(set, num)
		} else {
			set[num] = true
		}
	}

	return len(set) == 0
}

func divideArraySolution2(nums []int) bool {

	sort.Ints(nums)

	for i := 0; i < len(nums); i = i + 2 {

		if nums[i] != nums[i+1] {
			return false
		}
	}

	return true
}
