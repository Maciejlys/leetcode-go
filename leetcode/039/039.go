package leetcode

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

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
		for index := i; index < len(candidates); index++ {
			cur = append(cur, candidates[index])
			backtrack(index, sum+candidates[index], cur)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(0, 0, make([]int, 0))

	return result
}
