package leetcode

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	var getMaxArea func(i, j int, area *int)
	getMaxArea = func(i, j int, area *int) {
		if i < 0 || i == len(grid) {
			return
		}
		if j < 0 || j == len(grid[i]) {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		*area++

		getMaxArea(i+1, j, area)
		getMaxArea(i-1, j, area)
		getMaxArea(i, j+1, area)
		getMaxArea(i, j-1, area)
	}

	for i := range grid {
		for j := range grid[i] {
			area := 0
			getMaxArea(i, j, &area)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
