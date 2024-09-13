package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	amountAllowed := 0

	for i, fb := range flowerbed {
		left, mid, right := false, false, false

		if i-1 < 0 || flowerbed[i-1] == 0 {
			left = true
		}
		if i+1 >= len(flowerbed) || flowerbed[i+1] == 0 {
			right = true
		}
		if fb == 0 {
			mid = true
		}
		if left && mid && right {
			amountAllowed++
			flowerbed[i] = 1
		}
	}

	return n <= amountAllowed
}
