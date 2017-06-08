package leetcode

import "testing"

func testImpl(t *testing.T, in []int, expected int) {
	out := maxArea(in)
	t.Logf("in=%v, out=%d\n", in, out)

	if out != expected {
		t.Errorf("Got: %d, Expected: %d\n", out, expected)
	}
}

func TestMaxArea(t *testing.T) {
	testImpl(t, []int{1, 2, 3}, 2)
	testImpl(t, []int{0, 2, 3}, 2)
	testImpl(t, []int{1, 8, 2, 7, 3, 6, 4, 5}, 30)
}
