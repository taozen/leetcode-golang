package leetcode

/*
Given a roman numeral, convert it to an integer.
Input is guaranteed to be within the range from 1 to 3999.
*/

func romanToInt(s string) int {
	dict := make([]int, 128)
	dict['I'] = 1
	dict['V'] = 5
	dict['X'] = 10
	dict['L'] = 50
	dict['C'] = 100
	dict['D'] = 500
	dict['M'] = 1000

	slen := len(s)
	res := 0

	for i := 0; i < slen-1; i++ {
		if dict[s[i]] < dict[s[i+1]] {
			res -= dict[s[i]]
		} else {
			res += dict[s[i]]
		}
	}

	res += dict[s[slen-1]]
	return res
}
