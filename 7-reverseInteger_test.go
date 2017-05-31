package leetcode

import "testing"

func testImpl(t *testing.T, in, expected int) {
	out := reverse(in)
	t.Logf("in=%d, out=%d\n", in, out)
	if out != expected {
		t.Errorf("Got: %d, Expected: %d\n", out, expected)
	}
}

func TestReverseInteger(t *testing.T) {
	testImpl(t, 123, 321)
	testImpl(t, 320, 23)
	testImpl(t, 1534236469, 0)
}
