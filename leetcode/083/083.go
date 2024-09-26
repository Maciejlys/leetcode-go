package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	for curr := head; curr.Next != nil; {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
