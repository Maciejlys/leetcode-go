package leetcode

func permute(nums []int) [][]int {
	var res [][]int

	curr := make([]int, len(nums))
	visited := make([]bool, len(nums))

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			cp := make([]int, len(nums))
			copy(cp, curr)
			res = append(res, cp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if ok := visited[i]; ok == false {
				visited[i] = true
				curr[index] = nums[i]
				backtrack(index + 1)
				visited[i] = false
			}
		}
	}

	backtrack(0)

	return res
}
