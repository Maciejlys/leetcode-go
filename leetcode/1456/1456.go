package leetcode

func isVowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}

func maxVowels(s string, k int) int {
	maxV := 0
	curr := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			curr++
		}
	}

	maxV = curr

	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			curr--
		}
		if isVowel(s[i]) {
			curr++
		}
		maxV = max(maxV, curr)
	}

	return maxV
}
