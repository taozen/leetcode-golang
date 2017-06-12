package leetcode

// IntPair represents a <x, y> pair.
type IntPair struct {
	x, y int
}

// Triplet <x, y, z>
type Triplet struct {
	x, y, z int
}

func twoSum(nums []int, idx int) map[IntPair]bool {
	table := make(map[int]bool)
	ans := make(map[IntPair]bool)
	sum := -nums[idx]

	for i, v := range nums {
		if i == idx {
			continue
		}

		diff := sum - v
		if _, ok := table[diff]; ok {
			if v < diff {
				ans[IntPair{v, diff}] = true
			} else {
				ans[IntPair{diff, v}] = true
			}
		}
		table[v] = true
	}

	return ans
}

func threeSum(nums []int) [][]int {
	table := make(map[int]bool)
	result := make(map[Triplet]bool)

	for idx, v := range nums {
		if _, ok := table[v]; !ok {
			table[v] = true
			for pair := range twoSum(nums, idx) {
				if v < pair.x {
					result[Triplet{v, pair.x, pair.y}] = true
				} else if v > pair.y {
					result[Triplet{pair.x, pair.y, v}] = true
				} else {
					result[Triplet{pair.x, v, pair.y}] = true
				}
			}
		}
	}

	i, ans := 0, make([][]int, len(result))
	for key := range result {
		ans[i] = []int{key.x, key.y, key.z}
		i++
	}
	return ans
}
