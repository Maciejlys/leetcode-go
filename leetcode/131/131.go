package leetcode

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func partition(s string) [][]string {
	var result [][]string
	var curr []string

	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(s) {
			cp := make([]string, 0)
			cp = append(cp, curr...)
			result = append(result, cp)
			return
		}

		for i := index; i < len(s); i++ {
			subStr := s[index : i+1]
			if isPalindrome(subStr) {
				curr = append(curr, subStr)
				backtrack(i + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0)
	return result
}
