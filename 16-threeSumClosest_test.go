package leetcode

import "testing"

func testImpl(t *testing.T, in []int, target, expected int) {
	out := threeSumClosest(in, target)
	t.Logf("in=%v, out=%v", in, out)
	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestThreeSumClosest(t *testing.T) {
	testImpl(t, []int{-1, 2, 1, -4}, 1, 2)
	testImpl(t, []int{-1, 2, 1, -4}, 3, 2)
}
