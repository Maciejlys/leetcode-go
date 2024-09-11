package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	maxCandy := 0
	for _, n := range candies {
		if n > maxCandy {
			maxCandy = n
		}
	}

	for i := range candies {
		if candies[i]+extraCandies >= maxCandy {
			result[i] = true
		}
	}
	return result
}
