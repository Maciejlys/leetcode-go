package leetcode

type Loc struct {
	I int
	J int
}

var dirs = [4]Loc{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Empty  = 0
	Fresh  = 1
	Rotten = 2
)

type queue []*Loc

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) push(val *Loc) {
	*q = append(*q, val)
}

func (q *queue) shift() {
	*q = (*q)[1:]
}

func multiBfs(numFresh *int, q *queue, grid *[][]int) int {
	minutes := 0

	for !q.isEmpty() && *numFresh > 0 {
		numNodesToRemove := q.Len()

		for _, node := range *q {
			for _, dir := range dirs {
				i, j := node.I+dir.I, node.J+dir.J

				inBounds := i >= 0 && i < len(*grid) && j >= 0 && j < len((*grid)[i])
				if inBounds && (*grid)[i][j] == 1 {
					(*grid)[i][j] = 2
					*numFresh--

					q.push(&Loc{i, j})
				}

			}
		}

		for i := 0; i < numNodesToRemove; i++ {
			q.shift()
		}

		minutes++
	}

	return minutes
}

func orangesRotting(grid [][]int) int {
	numFresh := 0
	q := queue{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == Fresh {
				numFresh++
			} else if grid[i][j] == Rotten {
				q.push(&Loc{i, j})
			}
		}
	}

	minutes := multiBfs(&numFresh, &q, &grid)

	if numFresh > 0 {
		return -1
	}

	return minutes
}
