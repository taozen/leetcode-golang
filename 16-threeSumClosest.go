package leetcode

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	slen := len(nums)
	diff := (1 << 31) - 1

	for i := 0; i < slen-2; i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}

		lo, hi, sum := i+1, slen-1, target-nums[i]
		for lo < hi {
			twoSum := nums[lo] + nums[hi]

			if twoSum == sum {
				return target
			}

			newDiff := sum - twoSum

			if abs(newDiff) < abs(diff) {
				diff = newDiff
			}

			if twoSum < sum {
				lo++
			} else {
				hi--
			}
		}
	}

	return target - diff
}
