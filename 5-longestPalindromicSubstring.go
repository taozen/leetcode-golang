package leetcode

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}

		lo++
		hi--
	}

	return true
}

// O(N^3)
func longestPalindrome1(s string) string {
	totalLen := len(s)

	for len := totalLen; len > 0; len-- {
		for start := 0; start < totalLen-len+1; start++ {
			x := s[start : start+len]
			if isPalindrome(x) {
				return x
			}
		}
	}

	return ""
}

// O(N^2)
func longestPalindrome2(s string) string {
	n := len(s)

	if n == 0 {
		return ""
	}

	dp := make([][]bool, n)
	for idx := range dp {
		dp[idx] = make([]bool, n)
		dp[idx][idx] = true
	}

	maxPos, maxLen := 0, 1

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && (j-i < 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				slen := j + 1 - i
				if slen > maxLen {
					maxPos = i
					maxLen = slen
				}
			}
		}
	}

	return s[maxPos : maxPos+maxLen]
}

func expandPalindrome(s string, i, j int, pMaxPos, pMaxLen *int) {
	slen := len(s)
	for i >= 0 && j < slen && s[i] == s[j] {
		xlen := j - i + 1
		if xlen > *pMaxLen {
			*pMaxPos = i
			*pMaxLen = xlen
		}

		i--
		j++
	}
}

// O(N^2)
func longestPalindrome3(s string) string {
	maxPos, maxLen := 0, 0
	slen := len(s)

	for i := 0; i < slen; i++ {
		expandPalindrome(s, i, i, &maxPos, &maxLen)
		expandPalindrome(s, i, i+1, &maxPos, &maxLen)
	}

	return s[maxPos : maxPos+maxLen]
}

// O(N^2)
func longestPalindrome4(s string) string {
	ans := ""
	slen := len(s)

	for pivot := 0; pivot < slen; pivot++ {
		lo, hi := pivot, pivot

		// Move hi forward on repeated char such that
		// single-pivot and double-pivots cases are unified.
		for hi < slen-1 && s[hi] == s[hi+1] {
			hi++
		}
		pivot = hi

		for lo > 0 && hi < slen-1 && s[lo-1] == s[hi+1] {
			lo--
			hi++
		}

		xlen := hi - lo + 1
		if xlen > len(ans) {
			ans = s[lo : hi+1]
		}
	}

	return ans
}

func longestPalindrome(s string) string {
	return longestPalindrome4(s)
}
