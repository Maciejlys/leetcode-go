package leetcode

import "strings"

var vowels string = "aeiou"

func reverseVowels(s string) string {
	bytes := []byte(s)
	q := make([]byte, 0)

	for i, c := range bytes {
		if strings.Contains(vowels, strings.ToLower(string(c))) {
			q = append(q, c)
			bytes[i] = '*'
		}
	}

	for i, c := range bytes {
		if c == '*' {
			bytes[i] = q[len(q)-1]
			q = q[:len(q)-1]
		}
	}

	return string(bytes)
}
