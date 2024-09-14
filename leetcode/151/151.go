package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i := range len(words) >> 1 {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}

	return strings.Join(words, " ")
}
