package leetcode

/*
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target?
Find all unique quadruplets in the array which gives the sum of target.

Note: The solution set must not contain duplicate quadruplets.
For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	slen := len(nums)
	ans := [][]int{}

	for i := 0; i < slen-3; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		for j := i + 1; j < slen-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			lo, hi, sum := j+1, slen-1, target-nums[i]-nums[j]

			for lo < hi {
				twoSum := nums[lo] + nums[hi]

				if twoSum == sum {
					ans = append(ans, []int{nums[i], nums[j], nums[lo], nums[hi]})
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
	}

	return ans
}
