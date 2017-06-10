package leetcode

import "testing"

func testImpl(t *testing.T, s string, expected int) {
	out := romanToInt(s)
	t.Logf("in=%v, out=%v", s, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestRomanToInt(t *testing.T) {
	testImpl(t, "IV", 4)
	testImpl(t, "VII", 7)
	testImpl(t, "IX", 9)
	testImpl(t, "X", 10)
	testImpl(t, "XI", 11)
	testImpl(t, "XXXX", 40)
	testImpl(t, "MCMLIV", 1954)
	testImpl(t, "MCMXC", 1990)
	testImpl(t, "MMXIV", 2014)
}
