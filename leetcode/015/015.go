package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := range nums {
		num := nums[i]

		if num > 0 {
			break
		}

		if i > 0 && num == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {
			sum := num + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				left++
				right--

				for nums[left] == nums[left-1] && left < right {
					left++
				}
			}

			if sum < 0 {
				left++
			}

			if sum > 0 {
				right--
			}
		}

	}

	return res
}
