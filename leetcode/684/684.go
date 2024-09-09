package leetcode

func findRedundantConnection(edges [][]int) []int {
	parents := makeParents(len(edges))
	sizes := makeSizes(len(edges))

	var find func(int) int
	find = func(node int) int {
		if parents[node] != node {
			parents[node] = find(parents[node])
		}
		return parents[node]
	}

	union := func(node1, node2 int) bool {
		root1 := find(node1)
		root2 := find(node2)

		if root1 == root2 {
			return false
		}

		if sizes[root1] >= sizes[root2] {
			parents[root2] = root1
			sizes[root1] += sizes[root2]
		} else {
			parents[root1] = root2
			sizes[root2] += sizes[root1]
		}

		return true
	}

	for _, edge := range edges {
		node1, node2 := edge[0], edge[1]

		if !union(node1, node2) {
			return edge
		}
	}

	return []int{}
}

func makeParents(len int) map[int]int {
	tmp := make(map[int]int, len)

	for i := 0; i < len; i++ {
		tmp[i+1] = i + 1
	}

	return tmp
}

func makeSizes(len int) map[int]int {
	tmp := make(map[int]int, len)

	for i := 0; i < len; i++ {
		tmp[i+1] = 1
	}

	return tmp
}
