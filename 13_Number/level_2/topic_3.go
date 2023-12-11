package level2

// 判断 2的幂
// https://leetcode.cn/problems/power-of-two/description/

// 1000
// 0111
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
