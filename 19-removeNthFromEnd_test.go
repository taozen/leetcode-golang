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

func testImpl(t *testing.T, l *ListNode, n int, expected string) {
	in, out := mkString(l), mkString(removeNthFromEnd(l, n))
	t.Logf("in=%s, out=%s", in, out)

	if out != expected {
		t.Errorf("Got: %s, Expected: %s\n", out, expected)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	testImpl(t, makeList(0, 1, 2, 3, 4), 1, "[0 1 2 3]")
	testImpl(t, makeList(0, 1, 2, 3, 4), 2, "[0 1 2 4]")
	testImpl(t, makeList(0, 1, 2, 3, 4), 3, "[0 1 3 4]")
	testImpl(t, makeList(0, 1, 2, 3, 4), 4, "[0 2 3 4]")
	testImpl(t, makeList(0, 1, 2, 3, 4), 5, "[1 2 3 4]")
	testImpl(t, makeList(0), 1, "[]")
}
