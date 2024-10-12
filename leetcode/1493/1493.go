package leetcode

func longestSubarray(nums []int) int {
	l, r := 0, 0
	zeros := 0

	for ; r < len(nums); r++ {
		zeros += 1 - nums[r]
		if zeros > 1 {
			zeros -= (1 - nums[l])
			l++
		}
	}

	return r - l - 1
}
