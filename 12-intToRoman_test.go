package leetcode

import "testing"

func testImpl(t *testing.T, num int, expected string) {
	out := intToRoman(num)
	t.Logf("in=%v, out=%v", num, out)

	if out != expected {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestIntToRoman(t *testing.T) {
	testImpl(t, 1066, "MLXVI")
	testImpl(t, 3333, "MMMCCCXXXIII")
}
