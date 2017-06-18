package leetcode

/*
Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/

func strStr(s, sub string) int {
	slen, xlen := len(s), len(sub)
	if slen < xlen {
		return -1
	}

	if xlen == 0 {
		return 0
	}

	for i := 0; i <= slen-xlen; i++ {
		for j := 0; j < xlen; j++ {
			if s[i+j] != sub[j] {
				break
			}

			if j == xlen-1 {
				return i
			}
		}
	}

	return -1
}
