package leetcode

/*
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

For example, given array S = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

import "sort"

func threeSum(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	slen := len(nums)
	ans := [][]int{}

	for i := 0; i < slen-2; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		lo, hi, sum := i+1, slen-1, -nums[i]
		for lo < hi {
			twoSum := nums[lo] + nums[hi]

			if twoSum == sum {
				ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--

				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}

				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			} else if twoSum < sum {
				lo++
			} else {
				hi--
			}
		}
	}

	return ans
}
