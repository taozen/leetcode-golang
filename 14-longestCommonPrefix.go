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
