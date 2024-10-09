package leetcode

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	carry := 0

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}
	}

	if carry == 1 {
		return append([]int{1}, digits...)
	}

	return digits
}
