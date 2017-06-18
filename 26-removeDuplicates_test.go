package leetcode

import "testing"

func testImpl(t *testing.T, in []int, expected int) {
	t.Logf("in=%v", in)
	out := removeDuplicates(in)
	t.Logf("out=%d", out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestStrStr(t *testing.T) {
	testImpl(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, 8)
	testImpl(t, []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8}, 8)
}
