package level3

// 只出现一次的数字
// https://leetcode.cn/problems/single-number/description/
// 解法1 排序后找
// 解法2 set
// 解法3 位运算
func singleNumber(nums []int) (ans int) {
	for _, v := range nums {
		ans = ans ^ v
	}
	return ans
}
