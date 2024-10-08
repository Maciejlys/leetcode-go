package leetcode

func isPalindrome(s string) bool {

	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}

	return true
}

func longestPalindrome(s string) string {
	idx := make(map[rune][]int)

	for i, r := range s {
		idx[r] = append(idx[r], i)
	}

	longest := ""

	for _, v := range idx {
		for i := 0; i < len(v); i++ {
			for j := i; j < len(v); j++ {
				temp := s[v[i] : v[j]+1]
				if isPalindrome(temp) && len(temp) > len(longest) {
					longest = temp
				}
			}
		}
	}

	return longest
}
