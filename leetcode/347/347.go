package leetcode

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	occ := make(map[int]int)

	for i := range nums {
		occ[nums[i]]++
	}

	type kv struct {
		Key   int
		Value int
	}

	var ss []kv

	for k, v := range occ {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	res := []int{}

	for i, v := range ss {
		if i == k {
			break
		}
		res = append(res, v.Key)
	}

	return res
}
