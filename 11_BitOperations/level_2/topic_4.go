package level2

// 两整数之和
// https://leetcode.cn/problems/sum-of-two-integers/

func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a = a ^ b
		b = int(carry)
	}
	return a
}
