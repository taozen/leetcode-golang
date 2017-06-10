package leetcode

import "testing"

func testImpl(t *testing.T, ss []string, expected string) {
	out := longestCommonPrefix(ss)
	t.Logf("in=%v, out=%v", ss, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	testImpl(t, []string{"hello", "held", "helloworld"}, "hel")
	testImpl(t, []string{"abcdef", "bde"}, "")
	testImpl(t, []string{"abcdef", "abcd"}, "abcd")
}
