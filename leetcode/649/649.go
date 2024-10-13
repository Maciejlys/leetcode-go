package leetcode

import "bytes"

func predictPartyVictory(senate string) string {
	q := []byte(senate)
	bans := 0

	for bytes.ContainsRune(q, 'R') && bytes.ContainsRune(q, 'D') {
		curr := q[0]
		q = q[1:]

		if curr == 'R' {
			bans++
		} else {
			bans--
		}

		if curr == 'D' && bans < 0 || curr == 'R' && bans > 0 {
			q = append(q, curr)
		}
	}

	if q[0] == 'R' {
		return "Radiant"
	} else {
		return "Dire"
	}
}
