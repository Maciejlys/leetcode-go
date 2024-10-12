package leetcode

func longestOnes(nums []int, k int) int {
	longest, start, curr := 0, 0, 0

	for end := 0; end < len(nums); end++ {
		curr += nums[end]

		if end-start+1-curr <= k {
			longest = max(longest, end-start+1)
		} else {
			curr -= nums[start]
			start++
		}
	}

	return longest
}
