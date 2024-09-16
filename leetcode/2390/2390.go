package leetcode

func removeStars(s string) string {
	bytes := []byte(s)
	id := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] != '*' {
			bytes[id] = s[i]
			id++
		} else {
			id--
		}
	}
	return string(bytes[:id])
}
