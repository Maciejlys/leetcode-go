package leetcode

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		if set[num] == true {
			return true
		}
		set[num] = true
	}

	return false
}
