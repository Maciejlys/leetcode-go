package leetcode

type State struct {
	pacific  bool
	atlantic bool
}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	rows, cols := len(heights), len(heights[0])
	result := make([][]int, 0)
	visited := make([][]State, rows)
	for i := range visited {
		visited[i] = make([]State, cols)
	}

	var dfs func(i, j int, ocean string)
	dfs = func(i, j int, ocean string) {
		if ocean == "pacific" && visited[i][j].pacific {
			return
		}
		if ocean == "atlantic" && visited[i][j].atlantic {
			return
		}

		if ocean == "pacific" {
			visited[i][j].pacific = true
		}
		if ocean == "atlantic" {
			visited[i][j].atlantic = true
		}

		if visited[i][j].pacific && visited[i][j].atlantic {
			result = append(result, []int{i, j})
		}

		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			newI, newJ := i+dir[0], j+dir[1]
			if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && heights[newI][newJ] >= heights[i][j] {
				dfs(newI, newJ, ocean)
			}
		}
	}

	for i := 0; i < rows; i++ {
		dfs(i, 0, "pacific")
		dfs(i, cols-1, "atlantic")
	}
	for j := 0; j < cols; j++ {
		dfs(0, j, "pacific")
		dfs(rows-1, j, "atlantic")
	}

	return result
}
