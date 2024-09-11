package leetcode

func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int, 0)
	set := make(map[int]any, 0)

	for _, n := range arr {
		occ[n]++
	}

	for _, v := range occ {
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = 0
	}

	return true
}
