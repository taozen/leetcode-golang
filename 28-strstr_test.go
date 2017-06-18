package leetcode

import "testing"

func testImpl(t *testing.T, s, sub string, expected int) {
	out := strStr(s, sub)
	t.Logf("s=%s, sub=%s, out=%d", s, sub, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestStrStr(t *testing.T) {
	testImpl(t, "hello, world", "llo", 2)
	testImpl(t, "hello, world", "word", -1)
	testImpl(t, "abcde", "fg", -1)
	testImpl(t, "abcde", "", -1)
}
