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
	queue := []string{s}

	for len(queue) > 0 {
		head := queue[0]

		if isPalindrome(head) {
			return head
		}

		queue[0] = ""
		queue = queue[1:]
		queue = append(queue, head[:len(head)-1], head[1:])
	}

	return ""
}
