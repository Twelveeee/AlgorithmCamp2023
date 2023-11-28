package level2

// 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

func findMin(nums []int) int {
	low, high := 0, len(nums)-1

	for high > low {
		mid := (high + low) / 2
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}
