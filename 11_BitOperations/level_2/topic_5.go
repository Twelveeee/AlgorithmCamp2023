package level2

// 递归乘法
// https://leetcode.cn/problems/recursive-mulitply-lcci/description/

func multiply(A, B int) int {
	ans := 0
	// A == min
	if A > B {
		A, B = B, A
	}

	for i := 0; A > 0; i++ {
		if (A & 1) == 1 {
			ans += B << i
		}
		A >>= 1
	}
	return ans
}
