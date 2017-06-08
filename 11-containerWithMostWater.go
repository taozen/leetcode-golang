package leetcode

/*
Given n non-negative integers a1, a2, ..., an,
where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container,
such that the container contains the most water.
*/

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func maxArea(height []int) int {
	n, res := len(height), 0
	l, r := 0, n-1

	for l < r {
		res = max(res, (r-l)*min(height[l], height[r]))

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return res
}

func maxArea2(height []int) int {
	n, res := len(height), 0
	l, r := 0, n-1

	for l < r {
		h := min(height[l], height[r])
		res = max(res, (r-l)*h)

		for l < r && height[r] <= h {
			r--
		}

		for l < r && height[l] <= h {
			l++
		}
	}

	return res
}
