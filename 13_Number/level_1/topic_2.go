package level1

import "math"

// 整数反转
// https://leetcode.cn/problems/reverse-integer/description/

func reverse(x int) int {
	flag := 1
	ans := 0
	if x < 0 {
		flag = -1
		x = -x
	}
	for x > 0 {
		if ans >= (math.MaxInt32/10)+1 {
			return 0
		}
		ans = ans*10 + x%10

		x = x / 10
	}

	return ans * flag
}
