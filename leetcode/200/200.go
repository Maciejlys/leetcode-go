package leetcode

func numIslands(grid [][]byte) int {
	result := 0

	var visitConnected func(i, j int)
	visitConnected = func(i, j int) {
		if i < 0 || i == len(grid) {
			return
		}
		if j < 0 || j == len(grid[i]) {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'

		visitConnected(i+1, j)
		visitConnected(i-1, j)
		visitConnected(i, j+1)
		visitConnected(i, j-1)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				result++
				visitConnected(i, j)
			}
		}
	}

	return result
}
