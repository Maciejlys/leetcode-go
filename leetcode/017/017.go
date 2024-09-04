package leetcode

var letters = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) < 1 {
		return result
	}
	var backtrack func(index int, cur string)
	backtrack = func(index int, cur string) {
		if index == len(digits) {
			result = append(result, cur)
			return
		}
		key := digits[index]
		possible := letters[string(key)]

		for i := range possible {
			backtrack(index+1, cur+possible[i])
		}
	}

	backtrack(0, "")

	return result
}
