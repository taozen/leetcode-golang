package leetcode

/*
Given an array of integers, return indices of the two numbers such
that they add up to a specific target. You may assume that each input
would have exactly one solution, and you may not use the same element
twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	ans := make([]int, 2)

	for idx, val := range nums {
		if idx2, ok := table[val]; ok {
			return []int{idx2, idx}
		}

		table[target-val] = idx
	}

	return ans
}
