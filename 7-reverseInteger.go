package leetcode

/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

The input is assumed to be a 32-bit signed integer.
Your function should return 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	rx := 0

	for x != 0 {
		rem := x % 10
		val := rx*10 + rem

		if int32(val)/int32(10) != int32(rx) {
			return 0
		}

		rx = val
		x /= 10
	}

	return rx
}
