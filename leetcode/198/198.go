package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	cp := make([]int, len(nums))

	cp[0] = nums[0]
	cp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		cp[i] = max(cp[i-1], cp[i-2]+nums[i])

	}

	return cp[len(cp)-1]
}
