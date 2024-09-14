package leetcode

import "math"

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt32, math.MaxInt32

	for _, n := range nums {
		if n <= first {
			first = n
		} else if n < second {
			second = n
		} else {
			return true
		}
	}

	return false
}
