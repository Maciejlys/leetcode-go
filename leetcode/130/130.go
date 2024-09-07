package leetcode

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'E'
		dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := 0; i < rows; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][cols-1] == 'O' {
			dfs(i, cols-1)
		}
	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		if board[rows-1][j] == 'O' {
			dfs(rows-1, j)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'E' {
				board[i][j] = 'O'
			}
		}
	}
}
