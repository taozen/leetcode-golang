package leetcode

/*
Write a function to find the longest common prefix string amongst an array of strings.
*/

// Horizontal scaning.
func longestCommonPrefixHS(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	// This optimization looks over-engineering.
	minLen := len(ss[0])
	for _, s := range ss[1:] {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	ans := ss[0]
	for _, s := range ss[1:] {
		i := 0
		for i < minLen && i < len(ans) {
			if s[i] != ans[i] {
				break
			}
			i++
		}
		ans = ans[:i]
	}

	return ans
}

// Vertical scaning. It's faster in some cases, e.g. small len(ss) but large len(s).
func longestCommonPrefixVS(ss []string) string {
	if len(ss) == 0 {
		return ""
	}

	i, slen, ans := 0, len(ss[0]), ss[0]

outer:
	for i < slen {
		for _, s := range ss[1:] {
			if i == len(s) || s[i] != ans[i] {
				break outer
			}
		}
		i++
	}

	return ans[:i]
}

func longestCommonPrefix(ss []string) string {
	return longestCommonPrefixVS(ss)
}
