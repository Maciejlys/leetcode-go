package leetcode

func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber--
		result = string('A'+rune(columnNumber%26)) + result
		columnNumber /= 26
	}

	return result
}
