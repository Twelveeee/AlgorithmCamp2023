package level1

// 数组元素积的符号
// https://leetcode.cn/problems/sign-of-the-product-of-an-array/description/

func arraySign(nums []int) int {
	ans := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			ans *= -1
		}
	}
	return ans
}
