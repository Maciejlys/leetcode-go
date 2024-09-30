package leetcode

func majorityElement(nums []int) int {
	occ := make(map[int]int)

	for _, n := range nums {
		occ[n]++
	}

	for k, v := range occ {
		if v > len(nums)>>1 {
			return k
		}
	}

	return -1
}
