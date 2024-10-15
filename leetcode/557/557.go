package leetcode

import "strings"

func reverseWords(s string) string {
	fields := strings.Fields(s)

	for i := 0; i < len(fields); i++ {
		runes := []rune(fields[i])

		for j := 0; j < len(runes)>>1; j++ {
			runes[j], runes[len(runes)-1-j] = runes[len(runes)-1-j], runes[j]
		}
		fields[i] = string(runes)
	}

	return strings.Join(fields, " ")
}
