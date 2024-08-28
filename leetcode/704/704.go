package leetcode

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (right + left) >> 1
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
