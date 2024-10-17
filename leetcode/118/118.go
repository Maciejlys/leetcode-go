package leetcode

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	triangle := generate(numRows - 1)
	prevRow := triangle[len(triangle)-1]
	row := make([]int, 0)
	row = append(row, 1)

	for i := 1; i < len(prevRow); i++ {
		row = append(row, prevRow[i-1]+prevRow[i])
	}

	row = append(row, 1)
	triangle = append(triangle, row)

	return triangle
}
