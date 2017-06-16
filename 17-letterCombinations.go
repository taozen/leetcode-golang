package leetcode

import "bytes"

/*
Given a digit string, return all possible letter combinations that the number could represent.

Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*/

var dict = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans, aux := []string{""}, []string{}

	for i := range digits {
		d := digits[i]
		cs, ok := dict[d]
		if !ok {
			continue
		}

		aux = aux[:0]
		for _, s := range ans {
			buf := bytes.NewBuffer([]byte(s))
			for _, c := range cs {
				buf.WriteByte(c)
				aux = append(aux, buf.String())
				buf.Truncate(buf.Len() - 1)
			}
		}
		ans, aux = aux, ans
	}

	return ans
}
