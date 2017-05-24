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

func findMedianSortedArrays(num1, num2 []int) float64 {
	len1, len2 := len(num1), len(num2)
	totalLen := len1 + len2

	if totalLen == 0 {
		return 0.0
	}

	half := (totalLen - 1) / 2
	medianCnt := 1

	if totalLen%2 == 0 {
		medianCnt++
	}

	medianAry := make([]int, medianCnt)

	for i, i1, i2 := 0, 0, 0; i < half+medianCnt; i++ {
		idx := i % medianCnt
		if i1 == len1 {
			medianAry[idx] = num2[i2]
			i2++
		} else if i2 == len2 {
			medianAry[idx] = num1[i1]
			i1++
		} else if num1[i1] < num2[i2] {
			medianAry[idx] = num1[i1]
			i1++
		} else {
			medianAry[idx] = num2[i2]
			i2++
		}
	}

	if medianCnt == 1 {
		return float64(medianAry[0])
	}

	return (float64(medianAry[0]) + float64(medianAry[1])) / 2
}
