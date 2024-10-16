package leetcode

func findDisappearedNumbers(nums []int) []int {
	occ := make(map[int]int)
	result := make([]int, 0)

	for _, n := range nums {
		occ[n]++
	}

	for i := 1; i < len(nums); i++ {
		if _, ok := occ[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}
