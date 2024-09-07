package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int, (numCourses))
	seen := make(map[int]bool, (numCourses))

	for _, pre := range prerequisites {
		p1, p2 := pre[0], pre[1]

		preMap[p1] = append(preMap[p1], p2)
	}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if seen[course] == true {
			return false
		}
		if len(preMap[course]) == 0 {
			return true
		}

		seen[course] = true

		for _, c := range preMap[course] {
			if ok := dfs(c); !ok {
				return false
			}
		}

		seen[course] = false
		preMap[course] = []int{}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if ok := dfs(i); !ok {
			return false
		}
	}

	return true
}
