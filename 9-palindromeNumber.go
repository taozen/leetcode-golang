package leetcode

/*
Determine whether an integer is a palindrome. Do this without extra space.

Some hints:
Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem "Reverse Integer",
you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.
*/

import (
	"math"
)

// intuitive solution
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	digits, v := 0, x
	for v != 0 {
		v /= 10
		digits++
	}

	i, j, vl, vh := 0, digits-1, x, x
	for j > i {
		dl := vl % 10

		pow := int(math.Pow10(j))
		dh := vh / pow

		if dl != dh {
			return false
		}

		vh -= (dh * pow)
		vl /= 10
		i++
		j--
	}

	return true
}

// Reverse the lower half of the number and then
// compare it with the higher part.
func isPalindrome2(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	v := 0
	for x > v {
		v = v*10 + x%10
		x /= 10
	}

	return (x == v || x == v/10)
}

func isPalindrome(x int) bool {
	return isPalindrome2(x)
}
