package leetcode

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func iter(dividend, divisor, accN int) int {
	if dividend < divisor {
		return accN
	}

	acc, n := divisor, 1
	for (acc << 1) <= dividend {
		acc <<= 1
		n <<= 1
	}

	return iter(dividend-acc, divisor, accN+n)
}

const (
	maxInt = 2147483647
)

func divide(dividend, divisor int) int {
	if divisor == 0 || (dividend == -2147483648 && divisor == -1) {
		return maxInt
	}

	n := iter(abs(dividend), abs(divisor), 0)

	if dividend < 0 {
		n = -n
	}

	if divisor < 0 {
		n = -n
	}

	return n
}
