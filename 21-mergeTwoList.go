package leetcode

/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.
*/

// ListNode represents singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil {
		return l1
	}

	if l1 == nil {
		return l2
	}

	if l2.Val < l1.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
