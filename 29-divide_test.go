package leetcode

import "testing"

func testImpl(t *testing.T, dividend, divisor, expected int) {
	out := divide(dividend, divisor)
	t.Logf("a=%d, b=%d, a/b=%d", dividend, divisor, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestDivide(t *testing.T) {
	testImpl(t, 1024, 2, 1024/2)
	testImpl(t, 3388, 123, 3388/123)
}
