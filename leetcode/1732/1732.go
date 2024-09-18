package leetcode

func largestAltitude(gain []int) int {
	sum := 0
	largest := 0

	for _, n := range gain {
		sum += n
		largest = max(largest, sum)
	}

	return largest
}
