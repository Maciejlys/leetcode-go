package leetcode

func exist(board [][]byte, word string) bool {
	var traverse func(i int, j int, index int) bool
	traverse = func(i int, j int, index int) bool {
		if index == len(word) {
			return true
		}

		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] != word[index] {
			return false
		}

		preserve := board[i][j]
		board[i][j] = '.'

		if traverse(i+1, j, index+1) ||
			traverse(i-1, j, index+1) ||
			traverse(i, j+1, index+1) ||
			traverse(i, j-1, index+1) {
			return true
		}

		board[i][j] = preserve

		return false
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] && traverse(i, j, 0) {
				return true
			}
		}
	}

	return false
}
