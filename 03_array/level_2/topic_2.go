package level2

// 按奇偶排序数组
// https://leetcode.cn/problems/sort-array-by-parity/description/
// 解法1 新数组
// 解法2 头尾指针，尾指针为奇数数据
func sortArrayByParity(nums []int) []int {
	low, high := 0, len(nums)-1

	for low <= high {
		if nums[low]%2 == 1 {
			nums[low], nums[high] = nums[high], nums[low]
			high--
		} else {
			low++
		}
	}

	return nums
}
