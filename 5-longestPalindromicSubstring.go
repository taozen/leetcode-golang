package leetcode

/*
Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Input: "cbbd"
Output: "bb"
*/

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
				sLen := j + 1 - i
				if sLen > maxLen {
					maxPos = i
					maxLen = sLen
				}
			}
		}
	}

	return s[maxPos : maxPos+maxLen]
}

func expandPalindrome(s string, i, j int, pMaxPos, pMaxLen *int) {
	sLen := len(s)
	for i >= 0 && j < sLen && s[i] == s[j] {
		xLen := j - i + 1
		if xLen > *pMaxLen {
			*pMaxPos = i
			*pMaxLen = xLen
		}

		i--
		j++
	}
}

// O(N^2)
func longestPalindrome3(s string) string {
	maxPos, maxLen := 0, 0
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		expandPalindrome(s, i, i, &maxPos, &maxLen)
		expandPalindrome(s, i, i+1, &maxPos, &maxLen)
	}

	return s[maxPos : maxPos+maxLen]
}

// O(N^2)
func longestPalindrome4(s string) string {
	ans := ""
	sLen := len(s)

	for pivot := 0; pivot < sLen; pivot++ {
		lo, hi := pivot, pivot

		// Move hi forward on repeated char such that
		// single-pivot and double-pivots cases are unified.
		for hi < sLen-1 && s[hi] == s[hi+1] {
			hi++
		}
		pivot = hi

		for lo > 0 && hi < sLen-1 && s[lo-1] == s[hi+1] {
			lo--
			hi++
		}

		xLen := hi - lo + 1
		if xLen > len(ans) {
			ans = s[lo : hi+1]
		}
	}

	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// O(N) Manacher's algorightm
func longestPalindrome5(s string) string {
	sLen := len(s)

	// extended length that is equal to number of pivots
	xLen := 2*sLen + 1

	// array of palindrome length for each pivots
	palLenAry := make([]int, 0, xLen)

	// Using rightEdge as main loop iterator eliminates several corner cases.
	palLen, rightEdge := 0, 0

outer:
	for rightEdge < sLen {
		if rightEdge > palLen && s[rightEdge] == s[rightEdge-palLen-1] {
			palLen += 2
		} else {
			palLenAry = append(palLenAry, palLen)
			pivot := len(palLenAry) - 1

			// loop backwards
			start := pivot - 1
			end := pivot - palLen

			for i := start; i >= end; i-- {
				d := i - end
				if palLenAry[i] == d {
					palLen = d
					continue outer
				}
				palLenAry = append(palLenAry, min(d, palLenAry[i]))
			}

			palLen = 1
		}

		rightEdge++
	}

	// the last palindrome
	palLenAry = append(palLenAry, palLen)
	maxPalLen, maxPos := 0, 0

	for i := range palLenAry {
		if palLenAry[i] > maxPalLen {
			maxPos = i
			maxPalLen = palLenAry[i]
		}
	}

	start := (maxPos - maxPalLen + 1) / 2
	return s[start : start+maxPalLen]
}

func longestPalindrome(s string) string {
	return longestPalindrome5(s)
}
