package leetcode

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if s == "" {
		return true
	}

	left := 0

	for right := range t {
		if left > len(s)-1 {
			return true
		}
		if t[right] == s[left] {
			left++
		}
	}

	if left == len(s) {
		return true
	} else {
		return false
	}
}
