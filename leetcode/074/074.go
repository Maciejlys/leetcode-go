package leetcode

func flat(arr [][]int) []int {
	res := []int{}

	for i := range arr {
		for j := range arr[i] {
			res = append(res, arr[i][j])
		}
	}

	return res
}

func contains(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (right + left) >> 1
		if nums[middle] == target {
			return true
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	flatted := flat(matrix)
	return contains(flatted, target)
}
