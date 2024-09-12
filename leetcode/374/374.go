package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int
func guessNumber(n int) int {
	low := 0
	high := n
	for {
		mid := (high + low) >> 1
		result := guess(mid)

		if result == -1 {
			high = mid - 1
		}

		if result == 1 {
			low = mid + 1
		}

		if result == 0 {
			return mid
		}
	}
}
