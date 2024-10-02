package leetcode

import "math"

func isPerfectSquare(num int) bool {
	return math.Mod(math.Sqrt(float64(num)), 1) == 0
}
