package leetcode

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)

	for _, num := range nums1 {
		nums1Map[num] = 1
	}

	for _, num := range nums2 {
		nums2Map[num] = 1
		nums1Map[num]--
	}

	for _, num := range nums1 {
		nums2Map[num]--
	}

	res1 := make([]int, 0)

	for k, v := range nums1Map {
		if v > 0 {
			res1 = append(res1, k)
		}
	}

	res2 := make([]int, 0)

	for k, v := range nums2Map {
		if v > 0 {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}
