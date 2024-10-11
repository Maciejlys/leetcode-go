package leetcode

import "math"

func findPeakElement(nums []int) int {
	peak, index := math.MinInt32, 0

	for i, n := range nums {
		if n > peak {
			peak = n
			index = i
		}
	}

	return index
}
