package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
