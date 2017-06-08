package leetcode

func matchStar(c byte, s, p string) bool {
	//fmt.Printf("matchStar: c=%c, s=%s, p=%s\n", c, s, p)
	for i := 0; i <= len(s); i++ {
		if matchHere(s[i:], p) {
			return true
		}

		if i == len(s) || (s[i] != c && c != '.') {
			break
		}
	}

	return false
}

func matchHere(s string, p string) bool {
	//fmt.Printf("matchHere: s=%s, p=%s\n", s, p)
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) > 1 && p[1] == '*' {
		return matchStar(p[0], s, p[2:])
	}

	if len(s) == 0 {
		return false
	}

	if p[0] != '.' && p[0] != s[0] {
		return false
	}

	return matchHere(s[1:], p[1:])
}

// Dynamic Programing, O(m*n)
func isMatchDP(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)

	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for j := 1; j <= n; j++ {
		dp[0][j] = j > 1 && p[j-1] == '*' && dp[0][j-2]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				dp[i][j] = (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-2] || ((p[j-2] == s[i-1] || p[j-2] == '.') && dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}

func isMatch(s string, p string) bool {
	return isMatchDP(s, p)
}
