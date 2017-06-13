package leetcode

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
