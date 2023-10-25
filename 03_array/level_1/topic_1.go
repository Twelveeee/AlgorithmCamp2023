package level1

import "sort"

// 单调数列
// https://leetcode.cn/problems/monotonic-array/
func isMonotonic(nums []int) bool {
	return sort.IntsAreSorted(nums) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}

func isMonotonicV2(nums []int) bool {
	flagA, flagB := true, true

	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] < nums[i+1] {
			flagA = false
		}
		if nums[i] > nums[i+1] {
			flagB = false
		}

		if !flagA && !flagB {
			return false
		}
	}

	return flagA || flagB
}
