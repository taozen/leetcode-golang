package leetcode

import "testing"

func testImpl(t *testing.T, in, expected string) {
	out := longestPalindrome(in)
	t.Logf("in: %s, out: %s\n", in, out)

	if out != expected {
		t.Errorf("out: %s, expected: %s\n", out, expected)
	}
}

func TestLongestPalindrome(t *testing.T) {
	testImpl(t, "abccba", "abccba")
	testImpl(t, "abcdasdfghjkldcba", "a")
	testImpl(t, "aba", "aba")
	testImpl(t, "abbca", "bb")
	testImpl(t, "abbc", "bb")
	testImpl(t, "aaaabbbbaaaa", "aaaabbbbaaaa")
	testImpl(t, "bbbbbb", "bbbbbb")
	testImpl(t, "aabbaabbaabbaabc", "aabbaabbaabbaa")
}
