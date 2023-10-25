package level2

// 移除元素
// https://leetcode.cn/problems/remove-element/description/

// 解法1 快慢指针，慢指针为真实值，快指针为遍历值
// 解法2 头尾指针，尾指针为被删除的数据
func removeElement(nums []int, val int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		if nums[low] == val {
			for low <= high && nums[high] == val {
				high--
			}
			if high <= low {
				break
			}
			nums[low], nums[high] = nums[high], nums[low]
			high--
		}
		low++
	}
	return low
}
