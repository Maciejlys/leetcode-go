package leetcode

func singleNumber(nums []int) int {
	curr := 0

	for i := range nums {
		curr ^= nums[i]
	}

	return curr
}

func singleNumber2(nums []int) int {
	curr := 0

	for _, n := range nums {
		curr ^= n
	}

	return curr
}
