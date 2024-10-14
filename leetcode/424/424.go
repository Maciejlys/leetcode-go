package leetcode

func characterReplacement(s string, k int) int {
	count := [26]int{}
	res, maxW, left := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'A']++

		maxW = max(maxW, count[s[right]-'A'])

		if (right-left+1)-maxW > k {
			count[s[left]-'A'] -= 1
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}
