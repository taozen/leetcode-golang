package leetcode

import "testing"

func testImpl(t *testing.T, x, y []int, expected float64) {
	got := findMedianSortedArrays(x, y)
	t.Logf("a => %v, b => %v, median => %f\n", x, y, got)

	if got != expected {
		t.Errorf("Got: %f, Expected: %f\n", got, expected)
	}

	got = findMedianSortedArrays(y, x)
	t.Logf("a => %v, b => %v, median => %f\n", y, x, got)

	if got != expected {
		t.Errorf("Got: %f, Expected: %f\n", got, expected)
	}
}

func TestSingleInputEmpty(t *testing.T) {
	x := []int{}
	y := []int{1, 2, 3, 4, 5}
	testImpl(t, x, y, 3.0)
}

func TestBothInputEmpty(t *testing.T) {
	x := []int{}
	testImpl(t, x, x, 0.0)
}

func TestEqualLength(t *testing.T) {
	x := []int{1, 3, 5, 7, 9}
	y := []int{2, 4, 6, 8, 10}
	testImpl(t, x, y, 5.5)
}

func TestUnequalLength(t *testing.T) {
	x := []int{3, 5, 7}
	y := []int{2, 4, 7, 8, 10}
	testImpl(t, x, y, 6.0)
}

func TestUnequalLength2(t *testing.T) {
	x := []int{1, 2}
	y := []int{3, 4, 5}
	testImpl(t, x, y, 3.0)
}

func TestUnequalLength3(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{7, 8, 9}
	testImpl(t, x, y, 5.0)
}

func TestUnequalLength4(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{7, 8, 9, 10}
	testImpl(t, x, y, 5.5)
}
