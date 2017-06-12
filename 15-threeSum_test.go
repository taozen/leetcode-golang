package leetcode

import "testing"

func makeTriplet(a []int) Triplet {
	return Triplet{a[0], a[1], a[2]}
}

func convert(ary [][]int) map[Triplet]bool {
	ans := make(map[Triplet]bool)
	for _, v := range ary {
		ans[makeTriplet(v)] = true
	}
	return ans
}

func compareImpl(a [][]int, c map[Triplet]bool) bool {
	for _, v := range a {
		_, ok := c[makeTriplet(v)]
		if !ok {
			return false
		}
	}
	return true
}

func compare(a, b [][]int) bool {
	return len(a) == len(b) &&
		compareImpl(a, convert(b)) && compareImpl(b, convert(a))
}

func testImpl(t *testing.T, in []int, expected [][]int) {
	out := threeSum(in)
	t.Logf("in=%v, out=%v", in, out)
	if !compare(out, expected) {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestThreeSum(t *testing.T) {
	testImpl(t, []int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, 0, 1}, []int{-1, -1, 2}})
	testImpl(t, []int{-6, -4, -2, 0, 2, 4, 6}, [][]int{[]int{-6, 0, 6}, []int{-4, 0, 4}, []int{-2, 0, 2}, []int{-6, 2, 4}, []int{-4, -2, 6}})
	testImpl(t, []int{1, 2, 3}, [][]int{})
}
