package array

func partitionByPivot1(nums []int, pivot int) []int {
	lCount, gCount, pCount := 0, 0, 0

	for _, num := range nums {
		if num < pivot {
			lCount++
		} else if num > pivot {
			gCount++
		} else if num == pivot {
			pCount++
		}
	}

	i := 0
	j := lCount
	k := lCount + pCount

	result := make([]int, len(nums))
	for _, num := range nums {
		if num < pivot {
			result[i] = num
			i++
		} else if num > pivot {
			result[k] = num
			k++
		} else if num == pivot {
			result[j] = num
			j++
		}
	}

	return result
}

func partitionByPivot2(nums []int, pivot int) []int {

	size := len(nums)
	result := make([]int, len(nums))
	left, right := 0, size-1

	for i, j := 0, size-1; i < size; i, j = i+1, j-1 {
		if nums[i] < pivot {
			result[left] = nums[i]
			left++
		}

		if nums[j] > pivot {
			result[right] = nums[j]
			right--
		}
	}

	for left <= right {
		result[left] = pivot
		left++
	}

	return result
}
