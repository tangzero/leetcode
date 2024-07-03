package main

// @leet start
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

// @leet end

