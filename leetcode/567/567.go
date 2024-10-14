package leetcode

func equal(c1, c2 [26]int) bool {

	for i := 0; i < 26; i++ {
		if c1[i] != c2[i] {
			return false
		}
	}

	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	count1 := [26]int{}
	count2 := [26]int{}

	for i := 0; i < len(s1); i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if equal(count1, count2) {
			return true
		}

		count2[s2[i]-'a']--
		count2[s2[i+len(s1)]-'a']++
	}

	return equal(count1, count2)
}
