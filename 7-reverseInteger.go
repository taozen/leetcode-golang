package leetcode

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
