package level2

// 括号匹配问题
// https://leetcode.cn/problems/valid-parentheses/description/

// 解法1 栈
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0, 1000)

	for _, v := range []byte(s) {
		if pairs[v] > 0 {
			if len(stack) <= 0 || stack[len(stack)-1] != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}

	}

	return len(stack) == 0
}
