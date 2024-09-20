package leetcode

func tribonacci(n int) int {
	stack := []int{0, 1, 1}

	for i := 3; i <= n; i++ {
		stack = append(stack, stack[i-3]+stack[i-2]+stack[i-1])
	}

	return stack[n]
}
