package leetcode

type stack []int

func (this *stack) Pop() {
	if len(*this) > 0 {
		*this = (*this)[:len(*this)-1]
	}
}

func (this *stack) Push(val int) {
	*this = append(*this, val)
}

func (this *stack) Top() int {
	return (*this)[len(*this)-1]
}

func (this *stack) Len() int {
	return len(*this)
}

func asteroidCollision(asteroids []int) []int {
	stack := stack{asteroids[0]}
	for i := 1; i < len(asteroids); i++ {
		for stack.Len() > 0 && stack.Top() > 0 && stack.Top() < -asteroids[i] {
			stack.Pop()
		}
		if stack.Len() == 0 || stack.Top() < 0 || asteroids[i] > 0 {
			stack.Push(asteroids[i])
		}

		if stack.Top() == -asteroids[i] {
			stack.Pop()
		}
	}
	return stack
}
