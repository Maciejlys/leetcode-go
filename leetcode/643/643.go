package leetcode

func findMaxAverage(nums []int, k int) float64 {
	right := 0
	maxSum := 0

	for right < k {
		maxSum += nums[right]
		right++
	}

	sum := maxSum

	for right < len(nums) {
		sum += nums[right] - nums[right-k]
		maxSum = max(maxSum, sum)
		right++
	}

	return float64(maxSum) / float64(k)
}
