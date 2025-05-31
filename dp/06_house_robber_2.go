package dp

func rob2(nums []int) int {
	n := len(nums)

	first := []int{}
	second := []int{}

	for i := 0; i < n; i++ {
		if i != n-1 {
			first = append(first, nums[i])
		}
		if i != 0 {
			second = append(second, nums[i])
		}
	}

	return max(robMoney2Tab(first), robMoney2Tab(second))
}

func robMoney2Tab(nums []int) int {
	n := len(nums)

	prev1 := nums[0]
	prev2 := 0

	for i := 1; i < n; i++ {
		incl := prev2 + nums[i]
		excl := prev1

		ans := max(incl, excl)

		prev2 = prev1
		prev1 = ans
	}

	return prev1
}
