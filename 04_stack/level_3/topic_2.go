package level3

import "strconv"

// 逆波兰表达式
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, v := range tokens {
		val, err := strconv.Atoi(v)
		if err == nil {
			// 数字
			stack = append(stack, val)
		} else {
			// 操作
			num2, num1 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
