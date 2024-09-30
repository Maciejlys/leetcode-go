package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type ListNode = structures.ListNode

func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)

	for n := head; n != nil; n = n.Next {
		values = append(values, n.Val)
	}

	for left, right := 0, len(values)-1; left < right; left, right = left+1, right-1 {
		if values[left] != values[right] {
			return false
		}
	}

	return true
}
