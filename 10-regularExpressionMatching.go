package leetcode

import "fmt"

func matchStar(c byte, s, p string) bool {
	//fmt.Printf("matchStar: c=%c, s=%s, p=%s\n", c, s, p)
	for i := 0; i <= len(s); i++ {
		if isMatch(s[i:], p) {
			return true
		}

		if i == len(s) || (s[i] != c && c != '.') {
			break
		}
	}

	return false
}

func isMatch(s string, p string) bool {
	//fmt.Printf("isMatch: s=%s, p=%s\n", s, p)
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) > 1 && p[1] == '*' {
		return matchStar(p[0], s, p[2:])
	}

	if len(s) == 0 {
		return len(p) == 0
	}

	if p[0] != '.' && p[0] != s[0] {
		return false
	}

	return isMatch(s[1:], p[1:])
}
