package leetcode

func climbStairs(n int) int {
	last, preLast := 1, 1

	for range n - 1 {
		preLast, last = preLast+last, preLast
	}

	return preLast
}
