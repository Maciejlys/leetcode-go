package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0)
	preMap := make(map[int][]int, numCourses)
	seen := make(map[int]bool, numCourses)
	processed := make(map[int]bool, numCourses)

	for _, pre := range prerequisites {
		p1, p2 := pre[1], pre[0]
		preMap[p1] = append(preMap[p1], p2)
	}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if seen[course] {
			return false
		}
		if processed[course] {
			return true
		}

		seen[course] = true

		for _, c := range preMap[course] {
			if !dfs(c) {
				return false
			}
		}

		seen[course] = false
		processed[course] = true

		result = append(result, course)

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !processed[i] {
			if !dfs(i) {
				return []int{}
			}
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
