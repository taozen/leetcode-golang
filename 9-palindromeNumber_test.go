package leetcode

import "testing"

func testImpl(t *testing.T, in int, expected bool) {
	out := isPalindrome(in)
	t.Logf("in=%d, out=%v\n", in, out)
	if out != expected {
		t.Errorf("Got: %v, Expected: %v\n", out, expected)
	}
}

func TestIsPalindormeNumber(t *testing.T) {
	testImpl(t, 0, true)
	testImpl(t, 1, true)
	testImpl(t, 11, true)
	testImpl(t, 12, false)
	testImpl(t, 121, true)
	testImpl(t, 123, false)
	testImpl(t, 1221, true)
	testImpl(t, 1223, false)
	testImpl(t, 3123, false)
	testImpl(t, -12345, false)
	testImpl(t, -12321, false)
	testImpl(t, -1221, false)
	testImpl(t, -1211, false)
	testImpl(t, -2147447412, false)
}
