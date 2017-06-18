package leetcode

import "testing"

func compare(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func testImpl(t *testing.T, s string, expected []string) {
	out := letterCombinations(s)
	t.Logf("in=%v, out=%v", s, out)

	if !compare(out, expected) {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestLetterCombinations(t *testing.T) {
	testImpl(t, "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})
}
