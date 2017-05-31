package leetcode

import "testing"

func testImpl(t *testing.T, in string, expected int) {
	out := myAtoi(in)
	t.Logf("in=%s, out=%d\n", in, out)

	if out != expected {
		t.Errorf("Got: %d, Expected: %d\n", out, expected)
	}
}

func TestMyAtoi(t *testing.T) {
	testImpl(t, "123", 123)
	testImpl(t, "-123", -123)
	testImpl(t, "  -123", -123)
	testImpl(t, "  -0", 0)
	testImpl(t, " 0abc", 0)
}
