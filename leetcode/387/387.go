package leetcode

func firstUniqChar(s string) int {
	letters := [26]int{}

	for _, r := range s {
		letters[r-'a']++
	}

	for i, r := range s {
		if letters[r-'a'] == 1 {
			return i
		}
	}

	return -1
}
