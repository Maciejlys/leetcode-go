package leetcode

import "github.com/Maciejlys/leetcode-go/structures"

type ListNode = structures.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	merged := ListNode{}

	if list1.Val < list2.Val {
		merged = ListNode{Val: list1.Val, Next: mergeTwoLists(list1.Next, list2)}
	} else {
		merged = ListNode{Val: list2.Val, Next: mergeTwoLists(list1, list2.Next)}
	}

	return &merged
}
