package level2

import "strconv"

// 汇总区间
// https://leetcode.cn/problems/summary-ranges/description/
func summaryRanges(nums []int) []string {
	ans := make([]string, 0, 21)
	start, end := 0, 0

	for i := 0; i < len(nums); i++ {
		end = i
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == end {
				ans = append(ans, strconv.Itoa(nums[start]))
			} else {
				ans = append(ans, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[end]))
			}
			start = i + 1
		}
	}

	return ans
}
