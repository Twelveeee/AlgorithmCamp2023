package level3

// 计算器问题
// https://leetcode.cn/problems/basic-calculator-ii/description/
func calculate(s string) int {
	stack := make([]int, 0, 1023)

	preNum, ans := 0, 0
	preSign := '+'
	for i, ch := range s {
		isDigit := ch >= '0' && ch <= '9'

		if isDigit {
			preNum = preNum*10 + int(ch-'0')
		}

		if (!isDigit && ch != ' ') || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, preNum)
			case '-':
				stack = append(stack, -preNum)
			case '*':
				stack[len(stack)-1] *= preNum
			case '/':
				stack[len(stack)-1] /= preNum
			}

			preSign = ch
			preNum = 0
		}
	}

	for _, v := range stack {
		ans += v
	}
	return ans
}
