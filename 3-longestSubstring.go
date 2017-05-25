package leetcode

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	start := 0

	// golang doesn't have memset equivalent method,
	// so we use the zero value as an invalidation indicator.
	bucket := make([]int, 128)

	for idx, c := range s {
		start = max(start, bucket[c]+1)
		bucket[c] = idx + 1
		longest = max(longest, idx+2-start)
	}

	return longest
}
