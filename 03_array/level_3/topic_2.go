package level3

// 颜色分类
// https://leetcode.cn/problems/sort-colors/description/
// 解法1 双指针

func sortColors(nums []int) {
	low, high := 0, len(nums)-1
	i := 0
	for i <= high {
		if nums[i] == 0 {
			nums[i], nums[low] = nums[low], nums[i]
			low++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[high] = nums[high], nums[i]
			high--
		} else {
			i++
		}
	}

}
