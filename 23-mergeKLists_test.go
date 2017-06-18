package leetcode

import (
	"bytes"
	"fmt"
	"testing"
)

func makeList(digits ...int) *ListNode {
	sentinel := ListNode{0, nil}
	iter := &sentinel

	for _, i := range digits {
		iter.Next = &ListNode{i, nil}
		iter = iter.Next
	}

	return sentinel.Next
}

func mkString(l *ListNode) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for l != nil {
		buf.WriteString(fmt.Sprintf("%d ", l.Val))
		l = l.Next
	}

	if buf.Len() > 2 {
		buf.Truncate(buf.Len() - 1)
	}

	buf.WriteString("]")
	return buf.String()
}

func testImpl(t *testing.T, in []*ListNode, expected string) {
	out := mkString(mergeKLists(in))
	t.Logf("in=%v, out=%s", in, out)

	if out != expected {
		t.Errorf("Got: %s, Expected: %s\n", out, expected)
	}
}

func TestMergeKLists(t *testing.T) {
	testImpl(t,
		[]*ListNode{
			makeList(0, 5, 8, 11, 12, 16, 17),
			makeList(1, 3, 6, 9, 14, 15),
			makeList(2, 4, 7, 10, 13)},
		"[0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]")
}
