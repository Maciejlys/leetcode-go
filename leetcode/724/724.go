package leetcode

func pivotIndex(nums []int) int {
	left := 0
	right := 0

	for _, n := range nums {
		right += n
	}

	for i := range nums {
		if i > 0 {
			left += nums[i-1]
		}
		right -= nums[i]
		if left == right {
			return i
		}
	}

	return -1
}
