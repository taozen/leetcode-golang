package leetcode

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/

var pairs = map[byte]byte{')': '(', '}': '{', ']': '['}

func isValid(s string) bool {
	queue := []byte{}

	for i := range s {
		c := s[i]
		pair, ok := pairs[c]
		if !ok {
			queue = append(queue, c)
			continue
		}

		qlen := len(queue)
		if qlen == 0 || queue[qlen-1] != pair {
			return false
		}

		queue = queue[:qlen-1]
	}

	return len(queue) == 0
}
