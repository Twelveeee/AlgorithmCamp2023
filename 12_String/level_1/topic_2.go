package level1

import "math"

// 字符串转换整数 (atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/description/

func myAtoi(s string) int {
	i, ans, n := 0, 0, len(s)
	flag := 1
	for i < n && s[i] == ' ' {
		i++
	}
	if i < n && s[i] == '-' {
		flag = -1
		i++
	} else if i < n && s[i] == '+' {
		i++
	}

	for i < n {
		if s[i] < '0' || s[i] > '9' {
			return ans * flag
		}

		ans *= 10
		ans += int(s[i] - '0')
		if ans*flag < math.MinInt32 {
			return math.MinInt32
		} else if ans*flag > math.MaxInt32 {
			return math.MaxInt32
		}

		i++
	}

	return ans * flag
}
