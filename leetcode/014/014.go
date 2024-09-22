package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	smallest := 201

	for i := 1; i < len(strs); i++ {
		id := 0
		for j := 0; j < len(strs[i]); j++ {
			if j >= len(strs[0]) || strs[i][j] != strs[0][j] {
				break
			}
			id++
		}
		smallest = min(smallest, id)
	}

	return strs[0][:smallest]
}
