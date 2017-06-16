package leetcode

import "bytes"

/*
Given a digit string, return all possible letter combinations that the number could represent.

Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

var dict = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans, aux := []string{""}, []string{}

	for i := range digits {
		d := digits[i]

		if d < '2' || d > '9' {
			continue
		}

		aux = aux[:0]
		for _, s := range ans {
			buf := bytes.NewBuffer([]byte(s))
			cs := dict[d-'2']
			for i := range dict[d-'2'] {
				buf.WriteByte(cs[i])
				aux = append(aux, buf.String())
				buf.Truncate(buf.Len() - 1)
			}
		}
		ans, aux = aux, ans
	}

	return ans
}
