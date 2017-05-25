package leetcode

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func testImpl(t *testing.T, target int, input, expected []int) {
	output := twoSum(input, target)

	t.Logf("target: %d, input: %v, output: %v\n",
		target, input, output)

	if !equal(output, expected) {
		t.Errorf("Got: %v, Expected: %v", output, expected)
	}
}

func TestTwoSum(t *testing.T) {
	var input, expected []int
	var target int

	input = []int{2, 7, 11, 15}
	target = 9
	expected = []int{0, 1}
	testImpl(t, target, input, expected)

	input = []int{1, 1, 2, 3, 5, 8, 13, 21}
	target = 18
	expected = []int{4, 6}
	testImpl(t, target, input, expected)
}
