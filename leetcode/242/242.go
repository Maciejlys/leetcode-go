package leetcode

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, char := range s {
		chars[char] += 1
	}

	for _, char := range t {
		chars[char] -= 1
	}

	for _, occ := range chars {
		if occ != 0 {
			return false
		}
	}

	return true
}
