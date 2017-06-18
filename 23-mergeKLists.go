package leetcode

/*
Merge k sorted linked lists and return it as one sorted list.
*/

import "container/heap"

// ListNode represents a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap implements the heap interface for ListNode.
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Pop pops the minimum element from the heap.
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// O(NlogK) using MinHeap.
func mergeKLists(lists []*ListNode) *ListNode {
	sentinel := &ListNode{0, nil}
	tail := sentinel
	h := &NodeHeap{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	for h.Len() != 0 {
		tail.Next = heap.Pop(h).(*ListNode)
		tail = tail.Next

		if h.Len() > 0 && tail.Next != nil {
			heap.Push(h, tail.Next)
		}
	}

	return sentinel.Next
}
