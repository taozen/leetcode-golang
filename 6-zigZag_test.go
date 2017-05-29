package leetcode

import "testing"

func testImpl(t *testing.T, numRows int, input, expected string) {
	output := convert(input, numRows)
	t.Logf("input=%s, numRows=%d, output=%s\n", input, numRows, output)

	if output != expected {
		t.Errorf("Got: %s, Expected: %s\n", output, expected)
	}
}

func TestZigZag(t *testing.T) {
	testImpl(t, 3, "PAYPALISHIRING", "PAHNAPLSIIGYIR")
	testImpl(t, 1, "ABC", "ABC")
}
