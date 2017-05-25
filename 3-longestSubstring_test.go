package leetcode

import "testing"

func testImpl(t *testing.T, input string, expected int) {
	len := lengthOfLongestSubstring(input)
	t.Logf("%s => %d\n", input, len)
	if len != expected {
		t.Errorf("Got: %d, Expected: %d\n", len, expected)
	}
}

func TestLongestSubstring(t *testing.T) {
	testImpl(t, "abcabcbb", 3)
	testImpl(t, "bbbbb", 1)
	testImpl(t, "pwwkew", 3)
	testImpl(t, "pwwkewp", 4)
	testImpl(t, "aab", 2)
	testImpl(t, "dvdf", 3)

	s := "abcdefghijklmnopq"
	testImpl(t, s, len(s))
}
