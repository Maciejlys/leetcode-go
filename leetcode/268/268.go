package leetcode

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i, n := range nums {
		if i != n {
			return i
		}
	}

	return len(nums)
}
