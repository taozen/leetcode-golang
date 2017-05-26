package leetcode

/*
Find the median of the two sorted arrays. The overall run time complexity should be O(log(m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

Tao's comments:
Sorting an array costs O(NlogN) already. Why do we bother with an O(logN) algorithm
that finds the median? O(N) looks sufficient enough. Moreover, the leetcode OJ reveals
that the more complicated O(log(min(m, n))) algorithm gets a worse result than the
O(log(m+n)) one.

*/

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
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

//
// O(log(m+n)) solution
//
func findMedian1(x, y []int) float64 {
	totalLen := len(x) + len(y)

	if totalLen == 0 {
		return 0.0
	}

	m1 := findKthSmallest(totalLen/2+1, x, y)
	if totalLen%2 != 0 {
		return float64(m1)
	}

	m2 := findKthSmallest(totalLen/2, x, y)
	return float64(m1+m2) / 2.0
}

func median(x []int) float64 {
	len := len(x)
	return float64(x[(len-1)/2]+x[len/2]) / 2.0
}

//
// O(log(min(m,n))) solution
//
func findMedian2(x, y []int) float64 {
	lenX, lenY := len(x), len(y)

	if lenX+lenY == 0 {
		return 0.0
	}

	if lenX > lenY {
		x, y, lenX, lenY = y, x, lenY, lenX
	}

	if lenX == 0 {
		return median(y)
	}

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = ^MaxInt

	lo, hi := 0, 2*lenX
	for lo <= hi {
		cutX := lo + (hi-lo)/2
		cutY := lenX + lenY - cutX

		midLeftX := MinInt
		if cutX != 0 {
			midLeftX = x[(cutX-1)/2]
		}

		midRightX := MaxInt
		if cutX != 2*lenX {
			midRightX = x[cutX/2]
		}

		midLeftY := MinInt
		if cutY != 0 {
			midLeftY = y[(cutY-1)/2]
		}

		midRightY := MaxInt
		if cutY != 2*lenY {
			midRightY = y[cutY/2]
		}

		// fmt.Printf("lo=%d, hi=%d, cX=%d, cY=%d, L1=%d, R1=%d, L2=%d, R2=%d\n",
		// 	lo, hi, cutX, cutY, midLeftX, midRightX, midLeftY, midRightY)

		if midLeftX > midRightY {
			hi = cutX - 1
		} else if midRightX < midLeftY {
			lo = cutX + 1
		} else {
			val := max(midLeftX, midLeftY) + min(midRightX, midRightY)
			return float64(val) / 2.0
		}
	}

	return 0.0
}

func findMedianSortedArrays(x, y []int) float64 {
	return findMedian2(x, y)
}
