package leetcode

import (
	"strings"
	"unicode"
)

func parse(s string) string {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	return strings.Map(f, s)
}

func isPalindrome(s string) bool {
	parsed := parse(s)
	i := 0
	j := len(parsed) - 1

	for i < j {
		if parsed[i] != parsed[j] {
			return false
		}
		i++
		j--
	}

	return true
}
