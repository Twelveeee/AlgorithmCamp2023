package level2

// 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/

func hammingWeight(num uint32) int {
	ans := 0
	for num > 0 {
		num = num & (num - 1)
		ans++
	}
	return ans
}
