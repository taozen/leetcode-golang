package leetcode

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list. You may assume the two numbers
do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

// ListNode represents singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sentinel := ListNode{0, nil}
	iter := &sentinel

	for sentinel.Val != 0 || l1 != nil || l2 != nil {
		if l1 != nil {
			sentinel.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sentinel.Val += l2.Val
			l2 = l2.Next
		}

		iter.Next = &ListNode{sentinel.Val % 10, nil}
		iter = iter.Next
		sentinel.Val /= 10
	}

	return sentinel.Next
}
