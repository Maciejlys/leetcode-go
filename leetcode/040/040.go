package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)

	var backtrack func(i int, sum int, cur []int)
	backtrack = func(i int, sum int, cur []int) {
		if sum == target {
			cp := make([]int, len(cur))
			copy(cp, cur)
			result = append(result, cp)
			return
		}
		if sum > target {
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}
			cur = append(cur, candidates[j])
			backtrack(j+1, sum+candidates[j], cur)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0, 0, make([]int, 0))

	return result
}
