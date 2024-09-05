package leetcode

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	strBuilder := make([]byte, 0)
	str := []byte(strconv.Itoa(x))
	if str[0] == '-' {
		strBuilder = append(strBuilder, str[0])
		str = str[1:]
	}

	for i := len(str) - 1; i >= 0; i-- {
		strBuilder = append(strBuilder, str[i])
	}
	res, _ := strconv.Atoi(string(strBuilder))

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
