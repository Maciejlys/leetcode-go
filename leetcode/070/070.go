package leetcode

func climbStairs(n int) int {
	last, preLast := 1, 1

	for i := 1; i < n; i++ {
		preLast, last = preLast+last, preLast
	}

	return preLast
}
