package leetcode

func equalPairs(grid [][]int) int {
	pairs := 0
	visited := make(map[[200]int]int)
	keys := make([][200]int, 0)
	arr := [200]int{}

	for i, row := range grid {
		copy(arr[:], row)
		visited[arr]++

		column := make([]int, len(grid))
		for j := 0; j < len(grid); j++ {
			column[j] = grid[j][i]
		}
		copy(arr[:], column)
		keys = append(keys, arr)
	}

	for _, ar := range keys {
		if v, ok := visited[ar]; ok {
			pairs += v
		}
	}

	return pairs
}
