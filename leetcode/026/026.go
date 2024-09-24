package leetcode

func removeDuplicates(nums []int) int {
	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}

	return len(nums[:left+1])
}
