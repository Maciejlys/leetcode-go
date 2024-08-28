package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)
		anagrams[key] = append(anagrams[key], str)
	}

	result := [][]string{}

	for _, anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}
