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

func longestPalindrome(s string) string {
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
