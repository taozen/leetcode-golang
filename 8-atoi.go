package leetcode

/*
Implement atoi to convert a string to an integer.
Hint: Carefully consider all possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs).
You are responsible to gather all the input requirements up front.

**Requirements for atoi:
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found.
Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible,
and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number,
which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number,
or if no such sequence exists because either str is empty or it contains only whitespace characters,
no conversion is performed.

If no valid conversion could be performed, a zero value is returned.
If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.
*/

func myAtoi(s string) int {
	i, slen := 0, len(s)
	for i < slen && (s[i] == '\t' || s[i] == ' ') {
		i++
	}

	if i == slen {
		return 0
	}

	sign, ans, overflow := 1, 0, false

	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	for i < slen && s[i] >= '0' && s[i] <= '9' {
		val := ans*10 + int(s[i]-'0')
		if int32(val)/int32(10) != int32(ans) {
			overflow = true
			break
		}
		ans = val
		i++
	}

	if !overflow {
		return sign * ans
	}

	if sign == 1 {
		return 2147483647
	}

	return -2147483648
}
