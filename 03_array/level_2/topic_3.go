package level2

// 轮转数组
// https://leetcode.cn/problems/rotate-array/
// 解法1 临时数组
// 解法2 双指针交换
// 解法3 数组翻转，先整体翻转，
func rotate(nums []int, k int) {
	len := len(nums)
	newNums := make([]int, len)
	for i, v := range nums {
		newNums[(i+k)%len] = v
	}

	copy(nums, newNums)
}
