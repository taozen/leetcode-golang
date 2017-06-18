package leetcode

import "testing"

func compare(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	hash := map[string]bool{}
	for _, s := range x {
		hash[s] = true
	}

	for _, s := range y {
		if _, ok := hash[s]; !ok {
			return false
		}
	}

	return true
}

func testImpl(t *testing.T, n int, expected []string) {
	out := generateParenthesis(n)
	t.Logf("n=%d, out=%v", n, out)

	if !compare(out, expected) {
		t.Errorf("Got: %v, Expected: %v", out, expected)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	testImpl(t, 1, []string{"()"})
	testImpl(t, 2, []string{"()()", "(())"})
	testImpl(t, 3, []string{"()()()", "((()))", "(())()", "()(())", "(()())"})
}
