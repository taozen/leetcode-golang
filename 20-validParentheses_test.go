package leetcode

import "testing"

func testImpl(t *testing.T, s string, expected bool) {
	out := isValid(s)
	t.Logf("in=%v, out=%v", s, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestValidParentheses(t *testing.T) {
	testImpl(t, "()", true)
	testImpl(t, "[]", true)
	testImpl(t, "{}", true)
	testImpl(t, "[{}]", true)
	testImpl(t, "([{}])", true)
	testImpl(t, "([{}]", false)
	testImpl(t, "([{}]))", false)
	testImpl(t, ")([{}]))", false)
	testImpl(t, "}([{}]))", false)
	testImpl(t, "()}{}]))", false)
}
