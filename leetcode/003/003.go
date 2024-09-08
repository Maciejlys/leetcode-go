package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	bytes := []byte(s)
	left, maxLen := 0, 0

	for right := range bytes {
		for seen[bytes[right]] {
			seen[bytes[left]] = false
			left++
		}

		seen[bytes[right]] = true
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
