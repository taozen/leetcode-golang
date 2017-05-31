package leetcode

import "testing"

func testImpl(t *testing.T, in, pattern string, expected bool) {
	match := isMatch(in, pattern)
	t.Logf("in: %s, pattern: %s, match: %v\n", in, pattern, match)
	if match != expected {
		t.Errorf("Got: %v, Expected: %v\n", match, expected)
	}
}

func TestRegexpMatch(t *testing.T) {
	testImpl(t, "ab", ".*c", false)
	testImpl(t, "abc", "adc", false)
	testImpl(t, "abc", "a.c", true)
	testImpl(t, "abc", "a.*c", true)
	testImpl(t, "abc", ".*", true)
	testImpl(t, "abbccc", ".*b*c*ccc", true)
	testImpl(t, "abbccc", ".*b*c*c", true)
}
