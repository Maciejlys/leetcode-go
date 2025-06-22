package leetcode

func divideString(s string, k int, fill byte) []string {
	res := make([]string, 0)

	for i := 0; i < len(s); i = i + k {
		if i+k >= len(s) {
			partial := s[i:]
			for range k - len(partial) {
				partial = partial + string(fill)
			}
			res = append(res, partial)

			break
		}

		res = append(res, s[i:i+k])
	}

	return res
}
