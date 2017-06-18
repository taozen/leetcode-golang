package leetcode

import "testing"

func testImpl(t *testing.T, in []int, val, expected int) {
	t.Logf("in=%v, val=%d", in, val)
	out := removeElement(in, val)
	t.Logf("ary=%v, out=%d", in, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestStrStr(t *testing.T) {
	testImpl(t, []int{1}, 1, 0)
	testImpl(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, 9, 8)
	testImpl(t, []int{1, 1, 1, 1, 1}, 1, 0)
	testImpl(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 7)
	testImpl(t, []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8}, 8, 11)
}
