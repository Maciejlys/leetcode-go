package leetcode

func moveZeroes(nums []int) {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == 0 {
			nums = append(nums[:left], nums[left+1:]...)
			nums = append(nums, 0)
			right--
			left--
		}
		left++
	}
}
