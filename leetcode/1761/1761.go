package leetcode

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var strBuilder strings.Builder
	i := 0

	for i = range min(len(word1), len(word2)) {
		strBuilder.WriteByte(word1[i])
		strBuilder.WriteByte(word2[i])
	}

	strBuilder.WriteString(word1[i+1:])
	strBuilder.WriteString(word2[i+1:])

	return strBuilder.String()
}
