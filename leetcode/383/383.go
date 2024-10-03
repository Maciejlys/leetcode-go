package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)

	for _, s := range magazine {
		letters[s]++
	}

	for _, s := range ransomNote {
		letters[s]--
	}

	for _, v := range letters {
		if v < 0 {
			return false
		}
	}

	return true
}
