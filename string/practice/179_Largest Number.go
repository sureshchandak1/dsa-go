package practice

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	n := len(nums)
	elements := make([]string, n)

	for i := range n {
		elements[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(elements, func(i, j int) bool {
		first := elements[i] + elements[j]
		second := elements[j] + elements[i]
		return second < first
	})

	if elements[0] == "0" {
		return "0"
	}

	var res strings.Builder
	for _, val := range elements {
		res.WriteString(val)
	}

	return res.String()
}
