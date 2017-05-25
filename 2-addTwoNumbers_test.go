package leetcode

import (
	"bytes"
	"fmt"
	"math"
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

func mkInt(l *ListNode) int {
	val, idx := 0, 0
	for l != nil {
		val += l.Val * int(math.Pow10(idx))
		idx++
		l = l.Next
	}
	return val
}

func testImpl(t *testing.T, l1, l2 *ListNode) {
	lr := addTwoNumbers(l1, l2)
	s1 := mkString(l1)
	s2 := mkString(l2)
	sr := mkString(lr)

	t.Logf("addTwoNumbers %s %s => %s\n", s1, s2, sr)
	expected := mkInt(l1) + mkInt(l2)
	ans := mkInt(lr)

	if ans != expected {
		t.Errorf("Got: %d, Expected: %d\n", ans, expected)
	}
}

func TestZero(t *testing.T) {
	l1 := makeList(0)
	testImpl(t, l1, l1)
}

func TestEqualDigits(t *testing.T) {
	l1 := makeList(0, 1, 2, 3, 4)
	l2 := makeList(5, 6, 7, 8, 9)
	testImpl(t, l1, l2)
}

func TestUnequalDigits(t *testing.T) {
	l1 := makeList(2, 3, 4)
	l2 := makeList(5, 6, 7, 8, 9)
	testImpl(t, l1, l2)
}
