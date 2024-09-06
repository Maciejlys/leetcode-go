package util

import "sort"

func Sort2DInts(arr [][]int) {
	for _, subarray := range arr {
		sort.Ints(subarray)
	}

	sort.Slice(arr, func(i, j int) bool {
		for k := range arr[i] {
			if k >= len(arr[j]) || arr[i][k] != arr[j][k] {
				return arr[i][k] < arr[j][k]
			}
		}
		return len(arr[i]) < len(arr[j])
	})
}
