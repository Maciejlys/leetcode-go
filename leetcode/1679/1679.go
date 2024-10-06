package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	result := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			result++
			left++
			right--
		}

		if sum > k {
			right--
		}

		if sum < k {
			left++
		}
	}

	return result
}
