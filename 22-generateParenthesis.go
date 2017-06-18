package leetcode

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
[
	"((()))",
	"(()())",
	"(())()",
	"()(())",
	"()()()"
]
*/

// point represents a point in 2D coordinate.
type point struct {
	x, y int
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	ans, aux := []string{""}, []string{}
	states, statesAux := []point{point{0, 0}}, []point{}

	for states[0].y != n {
		for i := range ans {
			str, state := ans[i], states[i]

			if state.x != n {
				aux = append(aux, str+"(")
				statesAux = append(statesAux, point{state.x + 1, state.y})
			}

			if state.y < state.x {
				aux = append(aux, str+")")
				statesAux = append(statesAux, point{state.x, state.y + 1})
			}
		}

		ans, aux = aux, ans[:0]
		states, statesAux = statesAux, states[:0]
	}

	return ans
}
