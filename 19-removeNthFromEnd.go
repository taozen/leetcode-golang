package leetcode

/*
Given a linked list, remove the nth node from the end of list and return its head.

For example:
Given linked list: 1->2->3->4->5, and n = 2.
After removing the second node from the end, the linked list becomes 1->2->3->5.

Note:
Given n will always be valid.  Try to do this in one pass.
*/

// ListNode represents a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{0, head}
	node, mark := head, sentinel

	for node != nil {
		node = node.Next

		if n == 0 {
			mark = mark.Next
		} else {
			n--
		}
	}

	mark.Next = mark.Next.Next
	return sentinel.Next
}
