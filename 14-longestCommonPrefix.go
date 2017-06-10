package leetcode

/*
Write a function to find the longest common prefix string amongst an array of strings.
*/

func longestCommonPrefix(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	ans := ss[0]
	for _, s := range ss[1:] {
		i := 0
		for i < len(s) && i < len(ans) {
			if s[i] != ans[i] {
				break
			}
			i++
		}
		ans = ans[:i]
	}

	return ans
}

// A loop variant. It's faster in some cases, but TLE in leetcode OJ.
func longestCommonPrefix2(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	i, slen, ans := 0, len(ss[0]), ss[0]

outer:
	for i < slen {
		for _, s := range ss[1:] {
			if s[i] != ans[i] {
				break outer
			}
		}
	}

	return ans[:i]
}
