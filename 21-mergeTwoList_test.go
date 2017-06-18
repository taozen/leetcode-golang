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

func testImpl(t *testing.T, l1, l2 *ListNode, expected string) {
	in1, in2, out := mkString(l1), mkString(l2), mkString(mergeTwoLists(l1, l2))
	t.Logf("l1=%s, l2=%s, out=%s", in1, in2, out)

	if out != expected {
		t.Errorf("Got: %s, Expected: %s\n", out, expected)
	}
}

func TestMergeTwoLists(t *testing.T) {
	testImpl(t, makeList(0, 2, 4, 6, 8), makeList(1, 3, 5, 7, 9), "[0 1 2 3 4 5 6 7 8 9]")
}
