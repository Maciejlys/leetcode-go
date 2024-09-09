package leetcode

func minCostClimbingStairs(cost []int) int {
	for pos := len(cost) - 3; pos >= 0; pos-- {
		cost[pos] += min(cost[pos+1], cost[pos+2])
	}
	return min(cost[0], cost[1])
}
