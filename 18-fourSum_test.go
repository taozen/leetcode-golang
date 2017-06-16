package leetcode

import "testing"

type Quadruplet struct {
	x, y, z, q int
}

func makeQuadruplet(a []int) Quadruplet {
	return Quadruplet{a[0], a[1], a[2], a[3]}
}

func convert(ary [][]int) map[Quadruplet]bool {
	ans := make(map[Quadruplet]bool)
	for _, v := range ary {
		ans[makeQuadruplet(v)] = true
	}
	return ans
}

func compareImpl(a [][]int, c map[Quadruplet]bool) bool {
	for _, v := range a {
		_, ok := c[makeQuadruplet(v)]
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

func testImpl(t *testing.T, in []int, target int, expected [][]int) {
	out := fourSum(in, target)
	t.Logf("in=%v, out=%v", in, out)
	if !compare(out, expected) {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestFourSum(t *testing.T) {
	testImpl(t, []int{1, 0, -1, 0, -2, 2}, 0,
		[][]int{
			[]int{-1, 0, 0, 1},
			[]int{-2, -1, 1, 2},
			[]int{-2, 0, 0, 2},
		})
}
