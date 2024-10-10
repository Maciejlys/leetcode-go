package leetcode

import "math"

func scoreOfString(s string) int {
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		sum += int(math.Abs(float64(int(s[i]) - int(s[i+1]))))
	}
	return sum
}
