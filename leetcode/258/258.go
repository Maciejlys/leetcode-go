package leetcode

func addDigits(num int) int {
	sum := 0

	for num > 9 {
		sum += num % 10
		num /= 10
	}

	sum += num

	if sum > 9 {
		return addDigits(sum)
	}

	return sum
}
