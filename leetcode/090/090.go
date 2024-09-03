package leetcode

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {
	res := make(map[string][]int)
	var curr []int

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			cp := make([]int, len(curr))
			cp2 := make([]int, len(curr))
			copy(cp, curr)
			copy(cp2, curr)

			key := ""

			sort.Slice(cp2, func(i, j int) bool {
				return cp2[i] < cp2[j]
			})

			for i := range cp2 {
				key += strconv.Itoa(cp2[i])
			}

			if _, ok := res[key]; !ok {
				res[key] = cp
			}
			return
		}

		curr = append(curr, nums[index])
		backtrack(index + 1)
		curr = curr[:len(curr)-1]
		backtrack(index + 1)
	}

	backtrack(0)

	var values [][]int

	for k := range res {
		values = append(values, res[k])
	}

	return values
}
