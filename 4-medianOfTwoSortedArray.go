package leetcode

/*
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findKthSmallest(k int, x, y []int) int {
	if len(x) > len(y) {
		return findKthSmallest(k, y, x)
	}

	lenX := len(x)

	if lenX == 0 {
		return y[k-1]
	}

	if k == 1 {
		return min(x[0], y[0])
	}

	kX := min(k/2, lenX)
	kY := k - kX

	if x[kX-1] < y[kY-1] {
		return findKthSmallest(k-kX, x[kX:], y)
	}

	return findKthSmallest(k-kY, x, y[kY:])
}

func findMedianSortedArrays(x, y []int) float64 {
	totalLen := len(x) + len(y)

	if totalLen == 0 {
		return 0.0
	} else if totalLen%2 == 0 {
		m1 := findKthSmallest(totalLen/2, x, y)
		m2 := findKthSmallest(totalLen/2+1, x, y)
		return float64(m1+m2) / 2.0
	} else {
		return float64(findKthSmallest(totalLen/2+1, x, y))
	}
}
