package leetcode

func subsets(nums []int) [][]int {
	var result [][]int
	var curr []int

	var dfs func(int)
	dfs = func(index int) {
		if index == len(nums) {
			subset := make([]int, len(curr))
			copy(subset, curr)
			result = append(result, subset)
			return
		}
		curr = append(curr, nums[index])
		dfs(index + 1)
		curr = curr[:len(curr)-1] // Backtrack
		dfs(index + 1)
	}

	dfs(0)

	return result
}
